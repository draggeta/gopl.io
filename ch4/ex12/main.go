package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const baseURL = "https://xkcd.com"

// type comics struct {
// 	comic []comic
// }

type comic struct {
	Num        int
	URL        string `json:"img"`
	Title      string
	Transcript string
	Alt        string
}

func getComic(index int) (*comic, error) {
	var c comic
	var url string

	if index == 0 {
		url = fmt.Sprintf("%s/info.0.json", baseURL)
	} else {
		url = fmt.Sprintf("%s/%d/info.0.json", baseURL, index)
	}
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return &c, fmt.Errorf("HTTP request failed for comic %d: %s", index, err)
	}
	if resp.StatusCode != http.StatusOK {
		return &c, fmt.Errorf("Get failed for comic %d: %s", index, resp.Status)
	}
	json.NewDecoder(resp.Body).Decode(&c)

	return &c, nil
}

func getComicConcurrent(ch <-chan int, comics map[int]*comic, wg *sync.WaitGroup, ml *sync.RWMutex) {
	wg.Add(1)
	defer wg.Done()

	for i := range ch {
		c, err := getComic(i)
		if err != nil {
			fmt.Printf("Failed to download comic %d.\n%v\n", i, err)
			return
		}
		ml.Lock()
		comics[i] = c
		ml.Unlock()
		fmt.Printf("Downloaded comic %d\n", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func getComics(comics map[int]*comic) {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	var ml sync.RWMutex

	c, err := getComic(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lim := c.Num
	comics[lim] = c

	for i := 0; i < 5; i++ {
		go getComicConcurrent(ch, comics, &wg, &ml)
	}

	for i := 1; i <= lim; i++ {
		if _, ok := comics[i]; !ok {
			ch <- i
			continue
		}
		fmt.Printf("Skipped comic %d. Already downloaded.\n", i)
	}
	close(ch)

	wg.Wait()

	test, _ := json.Marshal(comics)
	ioutil.WriteFile("xkcddb.json", test, os.ModePerm)
}

func outComic(comics map[int]*comic, index int) {
	c, ok := comics[index]
	if ok {
		fmt.Printf("\nindex: %d\nurl: %s\ntranscript: %s\n", c.Num, c.URL, c.Transcript)
	} else {
		fmt.Printf("Index value %d doesn't exist", index)
	}
}

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Usage: xkcd (get|index|search)")
	}

	comics := make(map[int]*comic)
	text := make(map[string]map[int]bool)

	// Read the file
	content, err := ioutil.ReadFile("xkcddb.json")
	if err != nil {
		log.Fatalf("Error while reading a file %v", err)
	}

	err = json.Unmarshal(content, &comics)
	if err != nil {
		log.Fatalf("Error while unmarshal the content  %v", err)
	}

	if os.Args[1] == "get" {
		if len(os.Args) != 2 {
			log.Fatalln("Usage: xkcd get")
		}
		getComics(comics)
		return
	} else if os.Args[1] == "index" {
		if len(os.Args) != 3 {
			log.Fatalln("Usage: xkcd index <int>.")
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalln("Usage: xkcd index <int>.", err)
		}
		outComic(comics, index)
	} else if os.Args[1] == "search" {
		if len(os.Args) != 3 {
			log.Fatalln("Usage: xkcd search <str>.")
		}
		for k, v := range comics {
			input := v.Transcript
			scanner := bufio.NewScanner(strings.NewReader(input))
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				inner, ok := text[scanner.Text()]
				if !ok {
					inner = make(map[int]bool)
					text[scanner.Text()] = inner
				}
				inner[k] = true
			}
		}
		list := text[os.Args[2]]

		for index := range list {
			outComic(comics, index)
		}
	} else {
		log.Fatalln("Usage: xkcd (get|index|search)")
	}
}

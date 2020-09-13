package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type movie struct {
	Title  string
	Poster string
}

func getMovie(searchArgs []string) {
	var mov movie

	apiKey := os.Getenv("OMDB_API_KEY")
	unescapedSearchStr := strings.Join(searchArgs, " ")
	searchStr := url.QueryEscape(unescapedSearchStr)
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", apiKey, searchStr)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP request failed: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Get failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&mov); err != nil {
		fmt.Printf("JSON decode failed: %s", err)
	}

	presp, err := http.Get(mov.Poster)
	if err != nil {
		fmt.Printf("HTTP request failed: %s", err)
	}
	defer presp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Get failed: %s", presp.Status)
	}

	poster, err := ioutil.ReadAll(presp.Body)
	if err != nil {
		log.Fatal(err)
	}

	ext := path.Ext(mov.Poster)
	fileName := fmt.Sprintf("%s%s", mov.Title, ext)
	ioutil.WriteFile(fileName, poster, 0)
}

func main() {
	searchArgs := os.Args[1:]

	getMovie(searchArgs)
}

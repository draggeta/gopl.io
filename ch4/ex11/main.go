package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const apiURL = "https://api.github.com"

var actions = []string{"create", "read", "update", "close", "search"}

func create(repo, title string) {
	path := fmt.Sprintf("/repos/%s/issues", repo)
	s, err := textBody("")
	if err != nil {
		log.Fatalf("Failed to set the contents of the issue: %s", err)
	}

	body := Issue{Title: title, Body: s}
	issue, err := postPatch("POST", apiURL+path, body)
	if err != nil {
		log.Fatalf("Failed to create the issue: %s", err)
	}
	fmt.Printf("Number\t: %d\nUser\t: %s\nTitle\t: %s\nState\t: %s\nContent\t: %s\n", issue.Number, issue.User.Login, issue.Title, issue.State, issue.Body)

}

func read(repo, number string) {
	path := fmt.Sprintf("/repos/%s/issues/%s", repo, number)
	result, err := get(apiURL + path)
	if err != nil {
		log.Fatalf("Reading the issue failed: %s", err)
	}
	fmt.Printf("Number\t: %d\nUser\t: %s\nTitle\t: %s\nState\t: %s\nContent\t: %s\n", result.Number, result.User.Login, result.Title, result.State, result.Body)
}

func update(repo, number, title string) {
	rPath := fmt.Sprintf("/repos/%s/issues/%s", repo, number)
	result, err := get(apiURL + rPath)
	if err != nil {
		log.Fatalf("Reading the issue failed: %s", err)
	}

	path := fmt.Sprintf("/repos/%s/issues/%s", repo, number)
	s, err := textBody(result.Body)
	if err != nil {
		log.Fatalf("Failed to set the contents of the issue: %s", err)
	}

	body := Issue{Title: title, Body: s}
	issue, err := postPatch("PATCH", apiURL+path, body)
	if err != nil {
		log.Fatalf("Failed to create the issue: %s", err)
	}
	fmt.Printf("Number\t: %d\nUser\t: %s\nTitle\t: %s\nState\t: %s\nContent\t: %s\n", issue.Number, issue.User.Login, issue.Title, issue.State, issue.Body)
}

func close(repo, number string) {
	path := fmt.Sprintf("/repos/%s/issues/%s", repo, number)

	body := Issue{State: "closed"}
	issue, err := postPatch("PATCH", apiURL+path, body)
	if err != nil {
		log.Fatalf("Failed to create the issue: %s", err)
	}
	fmt.Printf("Number\t: %d\nUser\t: %s\nTitle\t: %s\nState\t: %s\nContent\t: %s\n", issue.Number, issue.User.Login, issue.Title, issue.State, issue.Body)
}

func textBody(c string) (string, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "ghtxt-")
	if err != nil {
		return "", fmt.Errorf("Couldn't create temporary file %s", err)
	}

	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(c)
	if err != nil {
		return "", fmt.Errorf("Couldn't set the file default content: %s", err)
	}
	tmpFile.Close()

	var editor string
	if runtime.GOOS == "windows" {
		editor = "notepad.exe"
	} else {
		editor = "vim"
	}

	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("cmd.Run() failed with %s", err)
	}

	text, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to read temp file %s\n %s", tmpFile.Name(), err)
	}
	return string(text), nil
}
func get(url string) (Issue, error) {
	var result Issue
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return result, fmt.Errorf("HTTP request failed: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("Get failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("JSON decode failed: %s", err)
	}

	return result, nil
}

func postPatch(method, url string, body Issue) (Issue, error) {
	var result Issue
	client := &http.Client{}

	data, err := json.Marshal(body)
	if err != nil {
		return result, fmt.Errorf("JSON marshaling failed: %s", err)
	}

	buf := bytes.NewBuffer(data)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return result, fmt.Errorf("HTTP request creation: %s", err)
	}
	req.SetBasicAuth(os.Getenv("GH_USER"), os.Getenv("GH_PASS"))

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return result, fmt.Errorf("HTTP request execution failed: %s", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return result, fmt.Errorf("Action failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("JSON decode failed: %s", err)
	}

	return result, nil
}

func main() {
	action := flag.String("a", "read", "Enter [create|read|update|close|search]")
	repo := flag.String("r", "", "Enter the owner/repository e.g.: golang/go")
	number := flag.String("n", "", "Enter the issue number")
	title := flag.String("t", "", "Enter the issue title")
	sterms := flag.String("s", "", "Search terms")

	flag.Parse()

	var actionCheck bool
	for _, s := range actions {
		if s == *action {
			actionCheck = true
			break
		}
	}
	if !actionCheck {
		fmt.Fprintln(os.Stderr, "Please enter an action: [create|read|update|close|search]")
		return
	}
	if *repo == "" {
		fmt.Fprintln(os.Stderr, "Please enter a repository")
		return
	}

	switch *action {
	case "create":
		create(*repo, *title)
	case "read":
		read(*repo, *number)
	case "update":
		update(*repo, *number, *title)
	case "close":
		close(*repo, *number)
	case "search":
		result, err := SearchIssues(strings.Split(*sterms, " "))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/

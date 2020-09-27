// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"html/template"
)

var issueCache IssuesSearchResult

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range $i, $v := .Items}}
<tr>
  <td><a href='/issue/{{$i}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='/user/{{$i}}'>{{.User.Login}}</a></td>
  <td><a href='/issue/{{$i}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var issue = template.Must(template.New("issue").Parse(`
<h1>{{.Number}} {{.Title}}</h1>
User:			{{ .User.Login }}<br>
Creation Date:	{{ .CreatedAt }}<br>
State:			{{ .State }}<br>
<p>
Issue:<br>
{{ .Body }}
</p>
`))

var user = template.Must(template.New("user").Parse(`
<h1>{{.Login}}</h1>
<img src="{{ .AvatarURL }}" alt="GitHub Avatar">
`))

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		if err := issueList.Execute(w, &issueCache); err != nil {
			log.Fatal(err)
		}
	}

	search := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		result, err := SearchIssues(r, &issueCache)
		if err != nil {
			log.Fatal(err)
		}
		if err := issueList.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	}

	issue := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/issue/"))
		if err != nil {
			log.Fatal(err)
		}
		if err := issue.Execute(w, issueCache.Items[id]); err != nil {
			log.Fatal(err)
		}
	}

	user := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/user/"))
		if err != nil {
			log.Fatal(err)
		}
		if err := user.Execute(w, issueCache.Items[id].User); err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/search/", search)
	http.HandleFunc("/issue/", issue)
	http.HandleFunc("/user/", user)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}

//!-

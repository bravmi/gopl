// Based on:
// https://github.com/torbiak/gopl/blob/master/ex4.14
//
// Usage: go run main.go bravmi gopl
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/bravmi/gopl/chap4/github"
)

type Cache struct {
	Issues         []github.Issue
	IssuesByNumber map[int]github.Issue
}

var issueListTemplate, issueTemplate *template.Template
var issueListPat, issuePat *regexp.Regexp

func init() {
	issueListTemplate = template.Must(template.ParseFiles("issueList.tpl"))
	issueTemplate = template.Must(template.ParseFiles("issue.tpl"))
	issueListPat = regexp.MustCompile(`^/issues/?$`)
	issuePat = regexp.MustCompile(`^/issues/(\d+)/?$`)
}

func cacheURL(issue *github.Issue) {
	issue.HTMLURL = fmt.Sprintf("/issues/%d", issue.Number)
}

func loadIssues(owner, repo string) (cache Cache, err error) {
	issues, err := github.GetIssues(owner, repo)
	if err != nil {
		return
	}
	for i := range issues {
		cacheURL(&issues[i])
	}
	cache.Issues = issues
	cache.IssuesByNumber = make(map[int]github.Issue, len(issues))
	for _, issue := range issues {
		cache.IssuesByNumber[issue.Number] = issue
	}
	return
}

func (cache Cache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if issueListPat.MatchString(r.URL.Path) {
		err := issueListTemplate.Execute(w, cache)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if ret := issuePat.FindStringSubmatch(r.URL.Path); ret != nil {
		s := ret[1]
		num, err := strconv.Atoi(s)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte("Invalid issue number"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		issue, ok := cache.IssuesByNumber[num]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte("Invalid issue number"))
			if err != nil {
				log.Fatal(err)
			}
		}
		err = issueTemplate.Execute(w, issue)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: OWNER REPO")
		os.Exit(1)
	}
	owner := os.Args[1]
	repo := os.Args[2]

	cache, err := loadIssues(owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", cache)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

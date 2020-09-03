// Issues prints a table of GitHub issues matching the search terms.
// usage: go run main.go repo:golang/go is:open json decoder
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bravmi/gopl/chap4/github"
)

type Issue = github.Issue

func filterIssues(issues []*Issue, pred func(*Issue) bool) []*Issue {
	res := []*Issue{}
	for _, item := range issues {
		if pred(item) {
			res = append(res, item)
		}
	}
	return res
}

func printIssues(issues []*Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	issues1 := filterIssues(result.Items, func(item *Issue) bool {
		return item.CreatedAt.After(time.Now().AddDate(0, -1, 0))
	})
	fmt.Printf("%d issues less than a month old:\n", len(issues1))
	printIssues(issues1)
	fmt.Println()

	issues2 := filterIssues(result.Items, func(item *Issue) bool {
		return item.CreatedAt.After(time.Now().AddDate(-1, 0, 0))
	})
	fmt.Printf("%d issues less than a year old:\n", len(issues2))
	printIssues(issues2)
	fmt.Println()

	issues3 := filterIssues(result.Items, func(item *Issue) bool {
		return item.CreatedAt.Before(time.Now().AddDate(-1, 0, 0))
	})
	fmt.Printf("%d issues more than a year old:\n", len(issues3))
	printIssues(issues3)
}

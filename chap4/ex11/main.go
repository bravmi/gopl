// Based on:
// https://github.com/torbiak/gopl/blob/master/ex4.11
//
// Usage:
// go run main.go search repo:golang/go is:open json decoder
// go run main.go create bravmi gopl
// go run main.go read bravmi gopl 4
// go run main.go edit bravmi gopl 4
// go run main.go close bravmi gopl 4
// go run main.go open bravmi gopl 4
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bravmi/gopl/chap4/github"
	"github.com/joho/godotenv"
)

func search(query []string) {
	result, err := github.SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func print(issue *github.Issue) {
	fmt.Println("issue:")
	b, err := json.MarshalIndent(issue, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func read(owner, repo, number string) {
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	print(issue)
}

func close_(owner, repo, number string) {
	issue, err := github.UpdateIssue(owner, repo, number, map[string]string{"state": "closed"})
	if err != nil {
		log.Fatal(err)
	}
	print(issue)
}

func open(owner string, repo string, number string) {
	issue, err := github.UpdateIssue(owner, repo, number, map[string]string{"state": "open"})
	if err != nil {
		log.Fatal(err)
	}
	print(issue)
}

func create(owner, repo, title string) {
	issue, err := github.CreateIssue(owner, repo, map[string]string{"title": title})
	if err != nil {
		log.Fatal(err)
	}
	print(issue)
}

func edit(owner string, repo string, number string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	tempfile, err := os.CreateTemp("", "edit_issue")
	if err != nil {
		log.Fatal(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(tempfile).Encode(map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	})
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tempfile.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	params := make(map[string]string)
	if err = json.NewDecoder(tempfile).Decode(&params); err != nil {
		log.Fatal(err)
	}

	issue, err = github.UpdateIssue(owner, repo, number, params)
	if err != nil {
		log.Fatal(err)
	}
	print(issue)
}

var usage string = `usage:
search QUERY
[read|edit|close|open] OWNER REPO ISSUE_NUMBER
create OWNER REPO
`

func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		if len(args) < 1 {
			usageDie()
		}
		search(args)
		return
	}

	if len(args) != 3 {
		usageDie()
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	if cmd == "create" {
		owner, repo, title := args[0], args[1], args[2]
		create(owner, repo, title)
		return
	}

	owner, repo, number := args[0], args[1], args[2]
	switch cmd {
	case "read":
		read(owner, repo, number)
	case "edit":
		edit(owner, repo, number)
	case "close":
		close_(owner, repo, number)
	case "open":
		open(owner, repo, number)
	}
}

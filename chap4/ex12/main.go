// Based on:
// https://github.com/vinceyuan/gopl-solutions/blob/master/ch04/ex4.12/ex4.12.go
//
// usage:
// go run main.go -from 1 -to 10
// go run main.go -search sketch
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const INDEX_FILE = "index.json"
const COMICS_URL = "https://xkcd.com/%d/info.0.json"

type Index struct {
	FilePath string
	Comics   map[int]*Comic
}

type Comic struct {
	Title      string
	Transcript string
	ImgUrl     string `json:"img"`
}

func getComic(url string) (*Comic, error) {
	log.Println("fetching", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}
	return &comic, nil
}

func getIndex(filePath string) *Index {
	index := &Index{
		FilePath: filePath,
		Comics:   make(map[int]*Comic),
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return index
	}
	if err := json.Unmarshal(data, &index.Comics); err != nil {
		log.Fatal(err)
	}
	return index
}

func (i *Index) save() {
	data, err := json.MarshalIndent(i.Comics, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(i.FilePath, data, 0644); err != nil {
		log.Fatal(err)
	}
}

func (i *Index) build(fromNum, toNum int) {
	for num := fromNum; num <= toNum; num++ {
		if _, exists := i.Comics[num]; exists {
			continue
		}
		url := fmt.Sprintf(COMICS_URL, num)
		comic, err := getComic(url)
		if err != nil {
			log.Println(err)
			continue
		}
		i.Comics[num] = comic
	}
}

func (i *Index) search(query string) []*Comic {
	var foundComics []*Comic
	query = strings.ToLower(query)
	for _, comic := range i.Comics {
		title := strings.ToLower(comic.Title)
		transcript := strings.ToLower(comic.Transcript)
		if strings.Contains(title, query) || strings.Contains(transcript, query) {
			foundComics = append(foundComics, comic)
		}
	}
	return foundComics
}

func (c *Comic) print() {
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return
	}
	log.Println(string(b))
}

func main() {
	fromNum := flag.Int("from", 0, "num from which to build index")
	toNum := flag.Int("to", 1, "num to which to build index")
	search := flag.String("search", "", "search for transcript and title")
	flag.Parse()

	index := getIndex(INDEX_FILE)
	if *fromNum != 0 && *toNum != 0 {
		index.build(*fromNum, *toNum)
		index.save()
	}
	if *search != "" {
		foundComics := index.search(*search)
		for _, comic := range foundComics {
			comic.print()
		}
	}
}

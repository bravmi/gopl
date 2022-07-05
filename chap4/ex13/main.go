// usage:
// go run main.go pay
// go run main.go 'pay it forward'

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const MOVIE_URL = "http://www.omdbapi.com/?apikey=%s&t=%s"

type Movie struct {
	Error    string
	Poster   string
	Response string
	Title    string
	Year     string
}

func getMovie(query string) (*Movie, error) {
	url := fmt.Sprintf(MOVIE_URL, os.Getenv("OMDB_API_KEY"), query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}
	return &movie, nil
}

func (m *Movie) print() {
	b, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return
	}
	log.Println(string(b))
}

func (m *Movie) downloadPoster() {
	log.Println("downloading", m.Poster)
	resp, err := http.Get(m.Poster)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("poster.jpg", data, 0644)
	if err != nil {
		return
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Usage: go run main.go query")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	query := url.QueryEscape(strings.Join(os.Args[1:], " "))
	movie, err := getMovie(query)
	if err != nil {
		log.Fatal(err)
	}
	if movie.Response != "True" {
		log.Fatal(movie.Error)
	}
	if movie.Poster == "" {
		log.Fatal("poster not found")
	}
	movie.print()
	movie.downloadPoster()
}

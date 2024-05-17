// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"io"
	"log"
	"net/http"
	"sort"
	"text/template"
	"time"
)

// !+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var tracksTemplate = template.Must(template.New("tracksHtml").Parse(`
<html>
	<head>
		<title>Tracks</title>
	</head>
	<body>
		<table>
			<tr>
				<th ><a href="?sort=Title">Title</a></th>
				<th ><a href="?sort=Artist">Artist</a></th>
				<th ><a href="?sort=Album">Album</a></th>
				<th ><a href="?sort=Year">Year</a></th>
				<th ><a href="?sort=Length">Length</a></th>
			</tr>
			{{- range . }}
			<tr>
				<td>{{ .Title }}</td>
				<td>{{ .Artist }}</td>
				<td>{{ .Album }}</td>
				<td>{{ .Year }}</td>
				<td>{{ .Length }}</td>
			</tr>
			{{- end }}
		</table>
	</body>
</html>
`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func printTracks(wr io.Writer, tracks []*Track) {
	err := tracksTemplate.Execute(wr, tracks)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort")
	sort.Stable(customSort{tracks, func(x, y *Track) bool {
		switch sortBy {
		case "Title":
			return x.Title < y.Title
		case "Artist":
			return x.Artist < y.Artist
		case "Album":
			return x.Album < y.Album
		case "Year":
			return x.Year < y.Year
		case "Length":
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(w, tracks)
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

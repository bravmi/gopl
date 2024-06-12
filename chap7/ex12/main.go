// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.

// usage:
// go run main.go
// http http://localhost:8000/list
// http http://localhost:8000/price\?item\=socks
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var listTemplate = template.Must(template.New("database").Parse(`
<html>
	<head>
		<title>Database</title>
	</head>
	<body>
		<table>
			<tr>
				<th>item</th>
				<th>price</th>
			</tr>
			{{- range $k, $v := . }}
			<tr>
				<td>{{ $k }}</td>
				<td>{{ $v }}</td>
			</tr>
			{{- end }}
		</table>
	</body>
</html>`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	err := listTemplate.Execute(w, db)
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// This file is a derivative work of "http4"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.12 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var list = template.Must(template.New("list").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <title>Track List</title>
  </head>
  <body>
    <table>
      <tr>
        <th>Item</th>
        <th>Price</th>
      </tr>
      {{range $item, $price := .}}
      <tr>
        <td>{{$item}}</td>
        <td>{{$price}}</td>
      </tr>
      {{end}}
    </table>
  </body>
</html>
`))


func (db database) list(w http.ResponseWriter, req *http.Request) {
	list.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

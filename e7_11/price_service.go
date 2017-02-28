// This file is a derivative work of "http4"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.11 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db :=  database{"shoes": 50, "socks": 5}
	lock := &sync.RWMutex{}

	http.Handle("/list", dbLockRLink(lock, http.HandlerFunc(db.list)))
	http.Handle("/price", dbLockRLink(lock, http.HandlerFunc(db.price)))
	http.Handle("/update", dbLockWLink(lock, http.HandlerFunc(db.update)))
	http.Handle("/remove", dbLockWLink(lock, http.HandlerFunc(db.remove)))
	http.Handle("/create", dbLockWLink(lock, http.HandlerFunc(db.create)))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		noSuchItemError(w, item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceParam := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceParam, 32)

	if err != nil || price < 0 {
		msg := fmt.Sprintf("invalid price: %s\n", priceParam)
		http.Error(w, msg, http.StatusBadRequest)
	} else if _, ok := db[item]; ok {
		db[item] = dollars(price)
		fmt.Fprintf(w, "price updated: %s\n", db[item])
	} else {
		noSuchItemError(w, item)
	}
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "%s removed\n", item)
	} else {
		noSuchItemError(w, item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceParam := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceParam, 32)

	if err != nil || price < 0 {
		msg := fmt.Sprintf("invalid price: %s\n", priceParam)
		http.Error(w, msg, http.StatusBadRequest)
	} else if len(item) == 0 {
		http.Error(w, "invalid item name\n", http.StatusBadRequest)
	} else if _, ok := db[item]; !ok {
		db[item] = dollars(price)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s created\n", item)
	} else {
		msg := fmt.Sprintf("item %s already exists\n", item)
		http.Error(w, msg, http.StatusBadRequest)
	}
}

func noSuchItemError(w http.ResponseWriter, item string) {
	msg := fmt.Sprintf("no such item: %q\n", item)
	http.Error(w, msg, http.StatusNotFound)
}

func dbLockRLink(lock *sync.RWMutex, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lock.RLock()
		defer lock.RUnlock()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func dbLockWLink(lock *sync.RWMutex, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lock.Lock()
		defer lock.Unlock()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// This file is a derivative work of the code found in page 126 of
// The Go Programming Language.
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.

// See exercise 5.5 of The Go Programming Language (http://www.gopl.io/)

// e5_5 prints the number of images and words in each HTML document
// fetched from the urls given as command line arguments.
package main

import (
	"fmt"
	"net/http"
	"os"

	"bufio"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			continue
		}

		fmt.Printf("url: %s, words: %d, images: %d\n", url, words, images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and return the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return countWordsAndImagesRec(n, nil)
}

func countWordsAndImagesRec(n *html.Node, p *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode &&
		(p == nil || (p.Data != "script" && p.Data != "style")) {
		words = countWords(n.Data)
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images = 1
	}

	wordsChild, imagesChild := countWordsAndImagesRec(n.FirstChild, n)
	wordsSibling, imagesSiblig := countWordsAndImagesRec(n.NextSibling, p)

	words += wordsChild + wordsSibling
	images += imagesChild + imagesSiblig

	return
}

func countWords(in string) (words int) {
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "scanning words: %s\n", err)
		os.Exit(1)
	}

	return
}

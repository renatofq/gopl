// This file is a derivative work of "findlinks3"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.13 of The Go Programming Language (http://www.gopl.io/)

// e5_13 crawls the web, starting with the URLs on the command line and
// saves local copies in the working dir, in a folder called out.
package main

import (
	"log"
	"os"

	"io"
	"net/http"
	urlib "net/url"

	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {

	fmt.Printf("url: %s\n", url)

	parsedURL, err := urlib.Parse(url)
	if err != nil {
		log.Print(err)
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return nil
	}
	defer resp.Body.Close()

	contentFile, err := saveFile(filepath.Join("out", parsedURL.Host),
		parsedURL.Path, resp.Body)
	if err != nil {
		log.Print(err)
		return nil
	}
	defer contentFile.Close()

	links, err := extract(contentFile)
	if err != nil {
		log.Print(err)
		return nil
	}

	var workList []string
	for _, l := range links {
		l = resp.Request.URL.ResolveReference(l)
		if err != nil {
			log.Print(err)
			continue
		}

		if resp.Request.URL.Host == parsedURL.Host && parsedURL.Host == l.Host {
			workList = append(workList, l.String())
		}
	}
	return workList
}

// saveFile creates a file on path and writes src's content to it.
// return the open file.
// FIXME: The path of a resource without a trailing slash will be used
// to name a file to store it. However, if this resource has subresouces
// saveFile will fail to create a directory to the subresource because
// the parent resource already got the name.
// A possible fix would be to walk all urls before saving the files and
// performing a topological sort. This is inefficent.
// Another fix could be always add a suffix to a downloaded file. But
// will not represent the tree structure nicely. A third possbility
// is to detect the occurrence and move the file to a temporary
// location, create a directory in it's place and move the file in
// with the name 'index'.
func saveFile(prefix, urlpath string, src io.Reader) (*os.File, error) {

	urlpath = evaluateDestinationPath(prefix, urlpath)
	dirname := filepath.Dir(urlpath)

	if err := os.MkdirAll(dirname, 0755); err != nil {
		return nil, fmt.Errorf("could not create directory %s: %v",
			dirname, err)
	}

	file, err := os.Create(urlpath)
	if err != nil {
		return nil, fmt.Errorf("could not create file %s: %v", urlpath, err)
	}

	_, err = io.Copy(file, src)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("fail to copy content to file %s: %v",
			urlpath, err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("fail to rewind file %s: %v", urlpath, err)
	}

	return file, nil
}

// evaluateDestinationPath evaluates the path of the destination file
func evaluateDestinationPath(prefix, urlpath string) string {
	if urlpath == "" {
		urlpath = "/"
	}

	if strings.HasSuffix(urlpath, "/") {
		// TODO: obtain file extension from content-type
		urlpath += "index"
	}

	urlpath = filepath.Join(prefix, urlpath)

	return filepath.FromSlash(urlpath)
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func extract(content io.Reader) ([]*urlib.URL, error) {

	doc, err := html.Parse(content)
	if err != nil {
		return nil, fmt.Errorf("parsing as HTML: %v", err)
	}

	var links []*urlib.URL
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := urlib.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link)
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

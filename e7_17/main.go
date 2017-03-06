// This file is a derivative work of "xmlselect"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.17 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []*selector // stack of element names

	sels, err := parseSelectors(os.Args[1:])
	if err != nil {
		log.Fatalf("xmlselect: %v\n", err)
	}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("xmlselect: %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			sel := &selector{
				name:  tok.Name.Local,
				class: findAttrValue(tok.Attr, "class"),
				id:    findAttrValue(tok.Attr, "id"),
			}
			stack = append(stack, sel) // push
			// log.Printf("push: %#v", sel)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if matchAll(stack, sels) {
				fmt.Printf("%v: %s\n", stack, tok)
			}
		}
	}
}

func findAttrValue(attrs []xml.Attr, name string) string {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return attr.Value
		}
	}

	return ""
}


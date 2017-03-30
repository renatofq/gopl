// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.18 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"encoding/xml"
	"io"
	"os"
	"fmt"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	var generateNodeTree func(*Element)
	decoder := xml.NewDecoder(os.Stdin)

	rootElem := &Element{}

	generateNodeTree = func (parent *Element) {
		for {
			token, err := decoder.Token()
			if err == io.EOF {
				return
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "xml fail: %v", err)
				os.Exit(1)
			}

			switch token := token.(type) {
			case xml.StartElement:
				e := &Element{
					Type: token.Name,
					Attr: token.Attr,
				}
				parent.Children = append(parent.Children, e)
				generateNodeTree(e)
			case xml.CharData:
				cd := CharData(token)
				parent.Children = append(parent.Children, cd)
			case xml.EndElement:
				return
			}
		}
	}

	generateNodeTree(rootElem)
	printElement(rootElem)

}

func printElement(e *Element) {
	var printElementRec func(*Element, string)

	printElementRec = func(e *Element, prefix string) {
		fmt.Printf("%s%s\n", prefix, e.Type.Local)
		prefix += "  "
		for _, node := range e.Children {
			switch node := node.(type) {
			case *Element:
				printElementRec(node, prefix)
			case CharData:
				fmt.Printf("%s%s", prefix, node)
			}
		}
	}

	printElementRec(e, "")
}

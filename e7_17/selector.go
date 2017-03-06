package main

import (
	"fmt"
	"bytes"
)

type selector struct {
	name  string
	class string
	id    string
}

// containsAll reports whether x contains the elements of y, in order.
func matchAll(x, y []*selector) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if match(x[0], y[0]) {
			y = y[1:]
		}

		x = x[1:]
	}
	return false
}

func match(x, y *selector) bool {
	// does not support multiple classes
	return (len(y.name) == 0 || x.name == y.name) &&
		(len(y.class) == 0 || x.class == y.class) &&
		(len(y.id) == 0 || x.id == y.id)
}

func parseSelectors(args []string) ([]*selector, error) {
	var sels []*selector

	for _, arg := range args {
		sel, err := parseSelector(arg)
		if err != nil {
			return nil, err
		}
		sels = append(sels, sel)
	}

	return sels, nil
}

func parseSelector(str string) (*selector, error) {

	if len(str) == 0 {
		return nil, fmt.Errorf("invalid selector: %s", str)
	}

	sel := &selector{}
	currSelRune := rune(0)
	lastIndex := 0
	for i, c := range str {
		if c == '.' || c == '#' {

			if i == 0 {
				currSelRune = c
				continue
			}

			if c == currSelRune {
				return nil, fmt.Errorf("invalid selector: %s", str)
			}

			selPart := str[lastIndex:i]
			if len(selPart) == 0 {
				return nil, fmt.Errorf("invalid selector: %s", str)
			}

			fillSelectorByType(sel, selPart, currSelRune)
			lastIndex = i

			currSelRune = c
		}
	}

	if lastIndex >= len(str) {
		return nil, fmt.Errorf("invalid selector: %s", str)
	}

	fillSelectorByType(sel, str[lastIndex:len(str)], currSelRune)

	return sel, nil
}

func fillSelectorByType(sel *selector, str string, selType rune) {
	switch selType {
	case 0: // Type selector
		sel.name = str
	case '.':
		sel.class = str[1:]
	case '#':
		sel.id = str[1:]
	default:
		panic("Invalid selector type")
	}
}

func (s *selector) String() string {
	buf := bytes.NewBufferString(s.name)
	if len(s.class) > 0 {
		buf.WriteRune('.')
		buf.WriteString(s.class)
	}
	if len(s.id) > 0 {
		buf.WriteRune('#')
		buf.WriteString(s.id)
	}
	return buf.String()
}

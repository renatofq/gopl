// This file is a derivative work of "sorting"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.8 of The Go Programming Language (http://www.gopl.io/)
package e7_8

import (
	"container/list"
	"fmt"
	"time"
)

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

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type trackSorter struct {
	tracks        []*Track
	fieldSortList *list.List
}

// Len length of the dataset
func (ts *trackSorter) Len() int {
	return len(ts.tracks)
}

// Swap swaps ith element int the dataset with the jth
func (ts *trackSorter) Swap(i, j int) {
	ts.tracks[i], ts.tracks[j] = ts.tracks[j], ts.tracks[i]
}

// Less returns true if the ith element is lesser than jth
func (ts *trackSorter) Less(i, j int) bool {
	for field := ts.fieldSortList.Front(); field != nil; field = field.Next() {

		switch field.Value.(string) {
		case "Title":
			if ts.tracks[i].Title == ts.tracks[j].Title {
				continue
			} else {
				return ts.tracks[i].Title < ts.tracks[j].Title
			}
		case "Artist":
			if ts.tracks[i].Artist == ts.tracks[j].Artist {
				continue
			} else {
				return ts.tracks[i].Artist < ts.tracks[j].Artist
			}
		case "Album":
			if ts.tracks[i].Album == ts.tracks[j].Album {
				continue
			} else {
				return ts.tracks[i].Album < ts.tracks[j].Album
			}
		case "Year":
			if ts.tracks[i].Year == ts.tracks[j].Year {
				continue
			} else {
				return ts.tracks[i].Year < ts.tracks[j].Year
			}
		case "Length":
			if ts.tracks[i].Length == ts.tracks[j].Length {
				continue
			} else {
				return ts.tracks[i].Length < ts.tracks[j].Length
			}
		default:
			panic(fmt.Sprintf("invalid field found in the field sort list: %s",
				field.Value))
		}
	}

	return false
}

func (ts *trackSorter) SortBy(fieldName string) {
	l := ts.fieldSortList
	for field := l.Front(); field != nil; field = field.Next() {
		if field.Value == fieldName {
			l.MoveToFront(field)
			break
		}
	}
}

func newTrackSorter(tracks []*Track) *trackSorter {
	ts := &trackSorter{tracks: tracks, fieldSortList: list.New()}
	ts.fieldSortList.PushBack("Title")
	ts.fieldSortList.PushBack("Artist")
	ts.fieldSortList.PushBack("Album")
	ts.fieldSortList.PushBack("Year")
	ts.fieldSortList.PushBack("Length")

	return ts
}

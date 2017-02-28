// This file is a derivative work of "sorting"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.8 of The Go Programming Language (http://www.gopl.io/)
package e7_8

import (
	"container/list"
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

type trackSort struct {
	tracks   []*Track
	sortList *list.List
}

type sortProperty struct {
	name string
	less func(x, y *Track) bool
}

// Len length of the dataset
func (ts *trackSort) Len() int {
	return len(ts.tracks)
}

// Swap swaps ith element int the dataset with the jth
func (ts *trackSort) Swap(i, j int) {
	ts.tracks[i], ts.tracks[j] = ts.tracks[j], ts.tracks[i]
}

// Less returns true if the ith element is lesser than jth
func (ts *trackSort) Less(i, j int) bool {
	for field := ts.sortList.Front(); field != nil; field = field.Next() {

		sortProp := field.Value.(*sortProperty)
		if !sortProp.less(ts.tracks[i], ts.tracks[j]) &&
			!sortProp.less(ts.tracks[j], ts.tracks[i]) {
			continue
		} else {
			return sortProp.less(ts.tracks[i], ts.tracks[j])
		}
	}

	return false
}

func (ts *trackSort) SortBy(fieldName string) {
	l := ts.sortList
	for field := l.Front(); field != nil; field = field.Next() {
		if prop := field.Value.(*sortProperty); prop.name == fieldName {
			l.MoveToFront(field)
			break
		}
	}
}

func newTrackSort(tracks []*Track) *trackSort {
	ts := &trackSort{tracks: tracks, sortList: list.New()}

	ts.sortList.PushBack(&sortProperty{
		name: "Title",
		less: func(x, y *Track) bool {
			return x.Title < y.Title
		},
	})

	ts.sortList.PushBack(&sortProperty{
		name: "Artist",
		less: func(x, y *Track) bool {
			return x.Artist < y.Artist
		},
	})

	ts.sortList.PushBack(&sortProperty{
		name: "Album",
		less: func(x, y *Track) bool {
			return x.Album < y.Album
		},
	})

	ts.sortList.PushBack(&sortProperty{
		name: "Year",
		less: func(x, y *Track) bool {
			return x.Year < y.Year
		},
	})

	ts.sortList.PushBack(&sortProperty{
		name: "Length",
		less: func(x, y *Track) bool {
			return x.Length < y.Length
		},
	})

	return ts
}

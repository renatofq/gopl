// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.9 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

var trackList = template.Must(template.New("tracklist").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <title>Track List</title>
  </head>
  <body>
    <h1>Track List</h1>
    <table>
      <tr>
        <th><a href='/?sortBy=Title'>Title</a></th>
        <th><a href='/?sortBy=Artist'>Artist</a></th>
        <th><a href='/?sortBy=Album'>Album</a></th>
        <th><a href='/?sortBy=Year'>Year</a></th>
        <th><a href='/?sortBy=Length'>Length</a></th>
      </tr>
      {{range .}}
      <tr>
        <td>{{.Title}}</td>
        <td>{{.Artist}}</td>
        <td>{{.Album}}</td>
        <td>{{.Year}}</td>
        <td>{{.Length}}</td>
      </tr>
      {{end}}
    </table>
  </body>
</html>
`))

type App struct {
	ts     *trackSorter
	tracks []*Track
}

func main() {
	app := &App{ts: newTrackSorter(tracks), tracks: tracks}
	http.HandleFunc("/", app.trackListHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// trackListHandler generates a html page containing the list of tracks
func (app *App) trackListHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if sortBy := r.Form.Get("sortBy"); sortBy != "" {
		app.ts.SortBy(sortBy)
		sort.Sort(app.ts)
	}
	trackList.Execute(w, app.tracks)
}

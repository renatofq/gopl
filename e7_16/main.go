// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.16 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/renatofq/gopl/e7_16/eval"
)

var index =`
<!DOCTYPE html>
<html>
  <head>
    <title>Calculator</title>
  </head>
  <body>
    <label for="expr">Expression:</label>
    <input id="expr" type="text" name="expr">
    <button onclick="eval()">Ok!</button> 
    <div id="ans">
    </div>
    <script type="text/javascript">
      document.getElementById('expr').onkeydown = function(e) {
        if(e.keyCode == 13){
          eval()
        }
      };

      function eval() {
        var xreq = new XMLHttpRequest();

        xreq.onreadystatechange = function() {
          if (xreq.readyState == XMLHttpRequest.DONE) {
              document.getElementById('ans').innerHTML = xreq.responseText;
          }
        };

        var q = '?expr=' + encodeURIComponent(document.getElementById('expr').value)
        xreq.open('GET', '/eval' + q, true);
        xreq.send();
      }
    </script>
  </body>
</html>
`

// indexHandler returns index page
func indexHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte(index))
}

// calcHandler evaluates the 'expr' query parameter and returns the result
func calcHandler(w http.ResponseWriter, r *http.Request) {
	e := r.URL.Query().Get("expr")
	if len(e) == 0 {
		http.Error(w, "No expression given", http.StatusBadRequest)
		return
	}

	expr, err := eval.Parse(e)
	if err != nil {
		msg := fmt.Sprintf("Invalid syntax: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	err = expr.Check()
	if err != nil {
		msg := fmt.Sprintf("Invalid expression: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%g", expr.Eval())
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/eval", calcHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

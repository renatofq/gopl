// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 5.16 of The Go Programming Language (http://www.gopl.io/)

// e5_16 exports Join that receives a separator and a varadic string
// and returns a joined version of it
import (
	"bytes"
)

// Join joins a number of arguments and returns
func Join(sep string, args ...string) string {

	if len(args) == 0 {
		return ""
	}

	buffer := bytes.NewBufferString(args[0])

	for _, s := range args[1:] {
		buffer.WriteString(sep)
		buffer.WriteString(s)
	}

	return buffer.String()
}

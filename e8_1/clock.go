// This file is a adapted of "clock2" to meet exercises needs
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 8.1 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
	"fmt"
)

var port = flag.Int("port", 8000, "Port to listen")

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()


	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}

}

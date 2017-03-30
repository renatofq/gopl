// This file is a derivative work of "reverb2"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 8.4 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c *net.TCPConn) {
	defer c.CloseWrite()

	var wg sync.WaitGroup

	func () {
		defer c.CloseRead()
		input := bufio.NewScanner(c)
		for input.Scan() {
			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				echo(c, s, 1*time.Second)
			}(input.Text())
		}
	}()

	wg.Wait()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn.(*net.TCPConn))
	}
}

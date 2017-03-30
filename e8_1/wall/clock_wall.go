// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 8.1 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
	"io"
)

type city struct {
	sync.Mutex
	name   string
	server string
	time   string
	err    error
	done   bool
}

func main() {
	cities := parseCityArgs()
	if len(cities) == 0 {
		log.Printf("No cities given\n")
		os.Exit(1)
	}

	for _, c := range cities {
		// uses shared memory because the book didn't brought
		// channels yet
		go updateCityTime(c)
	}

	ok := true
	for ok {
		ok = false
		for _, c := range cities {
			func() {
				c.Lock()
				defer c.Unlock()
				if !c.done {
					if c.err == nil {
						ok = true
						fmt.Printf("%s : %s\n", c.name, c.time)
					} else {
						log.Printf("%s is out: %v\n", c.name, c.err)
						c.done = true
					}
				}
			}()
		}

		time.Sleep(1 * time.Second)
	}
}

func updateCityTime(c *city) {
	defer func() {
		if msg := recover(); msg != nil {
			c.setError(fmt.Errorf("%s", msg))
		}
	}()

	log.Printf("connecting to %s\n", c.server)
	conn, err := net.Dial("tcp", c.server)
	if err != nil {
		log.Printf("fail to connect to %s\n", c.server)
		c.setError(err)
		return
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		log.Printf("scanned %s", c.name)
		c.setTime(scanner.Text())
	}

	if scanner.Err() != nil {
		c.setError(scanner.Err())
	} else {
		c.setError(io.EOF)
	}
}

func (c *city) setTime(time string) {
	c.Lock()
	defer c.Unlock()
	c.time = time
}

func (c *city) setError(err error) {
	c.Lock()
	defer c.Unlock()
	c.err = err
}

// parseCityArgs parse the command line arguments and returns citys to show
func parseCityArgs() []*city {
	var cities []*city
	for _, arg := range os.Args {
		strs := strings.Split(arg, "=")
		if len(strs) != 2 {
			fmt.Fprintf(os.Stderr, "ignoring invalid argument: %s\n", arg)
			continue
		}

		// don't care about input validation, is just an exercise anyway
		cities = append(cities, &city{name: strs[0], server: strs[1]})
	}

	return cities
}

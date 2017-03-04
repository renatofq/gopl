// Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 7.15 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/renatofq/gopl/e7_15/eval"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")

		if !scan.Scan() {
			break
		}

		line := strings.TrimSpace(scan.Text())
		if len(line) == 0 {
			continue
		}

		expr, err := eval.Parse(line)
		if err != nil {
			fmt.Printf("parse error: %v\n", err)
			continue
		}

		vars := make(map[eval.Var]bool)
		err = expr.Check(vars)
		if err != nil {
			fmt.Println(err)
			continue
		}

		env, err := readEnv(expr, vars)
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			continue
		}

		fmt.Printf("  %g\n", expr.Eval(env))
	}

	if err := scan.Err(); err != nil {
		fmt.Printf("reading error: %v\n", err)
	}
}

func readEnv(e eval.Expr, vars map[eval.Var]bool) (eval.Env, error) {
	var env = eval.Env{}

	scan := bufio.NewScanner(os.Stdin)
	for k := range vars {
		for {
			fmt.Printf("  %s := ", k)
			if !scan.Scan() {
				return nil, fmt.Errorf("incomplete environment")
			}

			v, err := strconv.ParseFloat(scan.Text(), 64)
			if err != nil {
				fmt.Println("Please inform a valid number")
				continue
			}

			env[k] = v
			break
		}
	}

	return env, nil
}

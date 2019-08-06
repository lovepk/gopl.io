// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

// sep和n变量分别是指向对应命令行标志参数变量的指针
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// go run main.go -n -s l a b c
func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	// 命令行参数-n使用的时候 *n为true，不使用的时候*n为false
	fmt.Println(*n)
	fmt.Println(*sep)
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("说明使用了命令行参数-n")
	}
}

//!-

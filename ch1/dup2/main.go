// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 注意countLines函数在其声明前被调用。函数和包级别的变量
// 可以任意顺序声明，并不影响其被调用。（译注：最好还是遵循一定的规范）
// map是一个由make函数创建的数据结构的引用。
// map作为为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本）
// 实际上指针是另一个指针了，但内部存的值指向同一块内存
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if (input.Text() == "end") { break }
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-

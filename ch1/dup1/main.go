// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// Scan函数在读到一行时返回true，不再有输入时返回false
	for input.Scan() {
		// 循环终止条件
		if input.Text() == "end" { break }
		counts[input.Text()]++
	}
	// map的迭代顺序并不确定，这种设计是有意为之的，因为map的key是没有规律变化的，不能依赖像
	// for循环这种依此变化的迭代顺序
	for line, n := range counts {
		if n > 1 {
			// 制表符\t和换行符\n
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-

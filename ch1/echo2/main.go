// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// 短变量声明,只能用在函数内部，而不能用于包变量
	s, sep := "", ""
	//  for循环的另一种形式，在某种数据类型的区间上遍历
	// 每次循环迭代，range产生一对值
	for _, arg := range os.Args[1:] {
		// 每次迭代字符串s的内容都会更新，如果数据量大，这种方式代价高昂
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-

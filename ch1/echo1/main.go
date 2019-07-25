// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"  // os以跨平台的方式提供和系统交互的函数和变量
)

func main() {
	// 变量会在声明时初始化，如果没有显示初始化，则被隐式初始化为其类型的零值。
	var s, sep string
	// 命令行参数即os包的Args变量，是一个字符串切片, 切片是长度动态变化的序列。
	// 切片通过s[i]访问单个元素，通过s[m:n]访问子序列。
	// 通过区间访问子序列时遵循左闭右开的原则，0 <= m <= n <= len(s),包含n-m个元素
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
//!-
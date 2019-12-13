// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 135.

// The squares program demonstrates a function value with state.
package main

import "fmt"

//!+
// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	// 对squares的一次调用会生成一个局部变量x并返回一个匿名函数
	// 每次调用匿名函数，都会使x的值加1
	return func() int {
		x++
		return x * x
	}
}

func main() {
	// 函数值记录了状态，它属于引用类型，所以不可以用==比较
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

//!-

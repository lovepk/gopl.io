// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 231.

// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import "fmt"

//!+
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

//  代码规范：当一个channel作为一个函数参数时，它一般总是专门用于只发送或者只接收
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	// close(in)  对一个只接收的channel调用close将是一个编译错误
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

//!-

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 229.

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// 当一个channel被关闭后，再向该channel发送数据将导致一个panic异常。接收者接收最后一个数据后，后续
	// 的操作将不再阻塞，会返回一个零值。所以接收者要判断channel是否关闭：
	// x, ok = <-naturals  通过布尔值ok判断，然后主动break跳出死循环
	//  for x := range naturals   会自己跳出循环
	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		//  squares是双向channel可以直接关闭
		close(squares)  
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

//!-

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import "crypto/sha256"

func main() {
	// 数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成

	//消息摘要有256bit大小，
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	// %x副词参数用于指定以十六进制的格式打印，%t用于打印布尔类型数据，%T用于显示一个值对应的数据类型
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// 消息摘要类型是[32]uint8,共32个元素，每个元素大小是8bit，总共大小是32*8 = 256bit
	fmt.Printf("%d", len(c1))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	//  一个slice由三个部分构成：指针，长度，容量
	// 指针指向 第一个slice元素对应的底层数组元素的地址，并不一定是底层数组首元素的地址 
	// 如果切片操作超出容量的上限将导致一个panic，但是超出长度则意味着扩展了slice，新slice的长度会变大
	// 合数组不同的是，slice之间不能比较，不能使用==操作符判断两个slice是否含有全部相等元素
	// 标准库提供了bytes.Equal来判断两个字节型slice是否相等，但对于其他类型的slice，我们必须自己展开每个元素进行比较

	// 安全的做法是直接禁止slice之间的比较操作，slice唯一合法的比较操作是和nil的比较
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev

//  s = []int{}  // len(s) == 0   s != nil
//  如果你需要测试一个slice是否是空的，使用len(s) == 0来判读，不应该使用s == nil来判断
// make([]T, len)
//  make([]T, len, cap)
//  在底层，make创建了一个匿名的数组变量，然后返回一个slice；第一种，slice是整个数组的view,第二个slice只引用了底层数组的前
// len个元素，但容量将包含整个数组

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 29.
//!+

// Boiling prints the boiling point of water.
package main

// 对于导入的包，则是对应源文件级的作用域，
// 因此只能在当前的文件中访问导入的fmt包，当前包的其它源文件无法访问在当前源文件导入的包
import "fmt"

// 一个声明语句将程序中的实体和一个名字关联，比如一个函数或一个变量。声明语句的作用域是指源代码中可以有效使用这个名字的范围
// 不要将作用域和生命周期混为一谈。声明语句的作用域对应的是一个源代码的文本区域；它是一个编译时的属性。
// 一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个运行时的概念
const boilingF = 212.0


// 句法块是由花括弧所包含的一系列语句,在代码中并未显式地使用花括号包裹起来，我们称之为词法块
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}

//!-

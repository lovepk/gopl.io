package main

import (
	"fmt"
)

func main () {
	// Go语言的数值类型包括几种不同大小的整数、浮点数和复数。每种数值类型都决定了对应的大小范围和是否支持正负符号
	// 例如int8  代表大小范围是-128到127 支持正负号

	// rune和int32是等价的类型，byte和uint8是等价类型   注意：int和int32也是不同的类型

	// 一个n-bit的有符号数的值域是从$-2^{n-1}$到$2^{n-1}-1$。无符号整数的所有bit位都用于表示非负数，值域是0到$2^n-1$
	// 例如，int8类型整数的值域是从-128到127，而uint8类型整数的值域是从0到255。

	// 对于int和uint 的大小  是根据当前设备的cpu确定的，一般是32或64bit

	// 超出的高位的bit位部分将被丢弃
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"

}
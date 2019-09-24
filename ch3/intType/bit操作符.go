
package main

import (
	"fmt"
)
/*

&      位运算 AND
|      位运算 OR
^      位运算 XOR  按位取反
&^     位清空 (AND NOT)  按位置零
<<     左移
>>     右移

*/
func main () {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\t%08b\t%08b\n", 1<<1, 1<<5, x) // "00000010"  "00100000" "00100010"
	fmt.Printf("%08b\t%08b\t%08b\n", 1<<1, 1<<2, y) // "00000010"  "00000100" "00000110"

	fmt.Printf("%08b\n", x&y)  // "00000010"
	fmt.Printf("%08b\n", x|y)  // "00100110"
	fmt.Printf("%08b\n", x^y)  // "00100100"
	fmt.Printf("%08b\t%08b\t%08b\n", x, ^y, x&^y) // "00100000"  功能和x&(^y)运算相同

	for i := uint(0); i < 8; i++ {
	    if x&(1<<i) != 0 { // membership test
	        fmt.Println(i) // "1", "5"
	    }
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}
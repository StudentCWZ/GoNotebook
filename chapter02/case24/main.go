/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 23:53
*/

package main

/*
	1.在官方的语言规范中，专门提到两个别名。 byte --> alias for uint8; rune alias for int32
	2. 别名类型无须转换，可直接赋值。
*/

func test(x byte) {
	println(x)
}

func main() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b

	test(c)
}


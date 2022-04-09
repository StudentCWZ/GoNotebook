/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 20:28
*/

package main

import "fmt"

/*
	1. 常量表示运行时恒定不可改变的值，通常是一些字面量。
	2. 使用常量就可用一个易于阅读理解的标识符号来代替 "魔术数学"，也使得在调整常量值时，无需修改所有引用代码。
	3. 常量值必须是编译期可确定的字符、字符串、数字或布尔值。
	4. 可指定常量类型，或由编译器通过初始值推断，不支持 C/C++ 数字类型后缀。
*/

const x, y int = 123, 0x22
const s = "hello, world!"
const c = '我' // rune(unicode code point)

const (
	i, f = 1, 0.123 // int, float64(默认)
	b    = false
)

func main() {
	fmt.Println(x, y, s, c, i, f, b)
}

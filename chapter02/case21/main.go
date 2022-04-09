/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 23:20
*/

package main

import (
	"fmt"
	"math"
)

/*
	1. 清晰完备的预定义基础类型，使得开发跨平台应用时无须过多考虑符号和长度差异。
	2. 支持八进制、十六进制以及科学记数法。标准库 math 定义了各数字类型的取值范围。
*/

func main() {
	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)
	fmt.Printf("0b%b, %#o,%#x\n", a, a, a)
	fmt.Println(math.MinInt8, math.MaxInt8)
}

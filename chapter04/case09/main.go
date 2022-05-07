/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 10:53
*/

package main

import (
	"errors"
	"log"
)

/*
	1. 标准库将 error 定义为接口类型，以便实现自定义错误类型。
	2. 按照惯例，error 总是最后一个返回参数。标准库提供了相关创建函数，可方便地创建包含简单错误文本的 error 对象。
*/

var errDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errDivByZero
	}
	return x / y, nil
}

func main() {
	z, err := div(5, 0)
	if err == errDivByZero {
		log.Fatalln(err)
	}
	println(z)
}

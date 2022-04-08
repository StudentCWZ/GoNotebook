/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:26
*/

package main

import (
	"errors"
	"fmt"
)

/*
	1. 函数可定义多个返回值，甚至对其命名。
*/

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	a, b := 10, 2       // 定义多个变量
	c, err := div(a, b) // 接收多返回值
	fmt.Println(c, err)
}

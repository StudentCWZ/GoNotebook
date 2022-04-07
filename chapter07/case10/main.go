/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 10:40
*/

package main

import "log"

type TestError struct{}

func (*TestError) Error() string {
	return "error"
}

func test(x int) (int, error) {
	// 错误
	//var err *TestError
	//if x < 0 {
	//	err = new(TestError)
	//	x = 0
	//} else {
	//	x += 100
	//}
	//return x, err

	// 修正
	if x < 0 {
		return 0, new(TestError)
	}
	return x + 100, nil
}

func main() {
	x, err := test(100)
	if err != nil {
		log.Fatalln("err != nil")
	}
	println(x)
}

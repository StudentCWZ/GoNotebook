/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 13:34
*/

package main

import "log"

/*
	1. 在延迟函数中 panic，不会影响延迟调用执行。而 recover 之后 panic，可被再次捕获。
	2. 另外，recover 必须在延迟调用函数中执行才能正常工作。
*/

func catch() {
	log.Println("catch:", recover())
}

func main() {
	defer catch()
	defer log.Println(recover())

	panic("i am dead")
}

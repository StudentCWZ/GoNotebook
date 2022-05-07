/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 11:39
*/

package main

import "log"

/*
	1. 连续调用 panic，仅最后一个会被 recover 捕获。
*/

func main() {
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("fatal")
			}
		}
	}()

	defer func() { // 类似重新抛出异常
		panic("you are dead")
	}()
	panic("i am dead")
}

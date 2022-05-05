/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/5 09:52
*/

package main

import (
	"errors"
	"log"
	"strconv"
)

/*
	1. 如须在多个条件块中使用局部变量，那么只能保留原层次，或直接使用外部变量。
	2. 对于某些过于复杂的组合条件，建议将其重构为函数。
	3. 函数调用虽然有一些性能损失，可却让主流程变得更加清爽。况且，条件语句独立之后，更易于测试，通常会改善代码可维护性。
*/

func check(s string) error {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil || n < 0 || n > 10 || n%2 != 0 {
		return errors.New("invalid number")
	}
	return nil
}

func ifCaseOne() {
	s := "9"

	n, err := strconv.ParseInt(s, 10, 64) // 使用外部变量

	if err != nil {
		log.Fatalln(err)
	} else if n < 0 || n > 10 { // 也可以拆分成另一个独立 if 块
		log.Fatalln("Invalid number ...")
	}
	println(n) // 避免 if 局部变量将该逻辑放在 else 块
}

func ifCaseTwo() {
	s := "9"

	if n, err := strconv.ParseInt(s, 10, 64); err != nil || n < 0 || n > 10 {
		log.Fatalln("Invalid number ...")
	}
	println("ok")
}

func ifCaseThree() {
	s := "9"

	if err := check(s); err != nil {
		log.Fatalln(err)
	}
	println("ok")
}

func main() {
	ifCaseOne()
	ifCaseTwo()
	ifCaseThree()
}

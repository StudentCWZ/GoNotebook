/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/6 10:02
*/

package main

import (
	"fmt"
	"log"
	"time"
)

/*
	1. 要实现传出参数(out)，通常建议使用返回值。当然，也可以继续用二级指针。
	2. 如果函数参数过多，建议将其重构为一个复合结构类型，也算是变相实现可选参数和命名实参功能。
*/

type serverOption struct {
	address string
	port    int
	path    string
	timeout time.Duration
	log     *log.Logger
}

func test(p **int) {
	x := 100
	*p = &x
}

func newOption() *serverOption { // 将过多参数独立成 option struct，既便于扩展参数类集，也方便通过 newOption 函数设置默认配置
	return &serverOption{ // 默认参数
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/test",
		timeout: time.Second * 5,
		log:     nil,
	}
}

func server(option *serverOption) {
	fmt.Printf("%#v\n", *option)
}

func main() {
	var p *int
	test(&p)
	println(*p)

	opt := newOption()
	opt.port = 8085 // 命名参数设置
	server(opt)
}

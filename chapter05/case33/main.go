package main

import (
	"fmt"
	"reflect"
)

/*
	1. 字段标签（tag）并不是注释，而是用来对字段进行描述的元数据。尽管它不属于数据成员，但却是类型的组成部分。
	2. 在运行期，可用反射获取标签信息。它常被用作格式校验，数据库关系映射等。
*/

type user struct {
	name string `昵称`
	sex  byte   `性别`
}

func main() {
	u := user{"Tom", 1}
	v := reflect.ValueOf(u)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}
}

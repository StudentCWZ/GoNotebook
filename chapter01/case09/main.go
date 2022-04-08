/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:48
*/

package main

import "fmt"

/*
	1. 将字典(map)类型内置，可直接从运行时层面获得性能优化。
*/

func main() {
	m := make(map[string]int) // 创建字典类型对象

	m["a"] = 1      // 添加或设置
	x, ok := m["b"] // 使用 ok-idiom 获取值，可知道 key/value 是否存在
	fmt.Println(x, ok)
	delete(m, "a")
}

/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 11:04
*/

package main

import "fmt"

/*
	1. 除 new/make 函数外，也可使用初始化表达式，编译器生成的指令基本相同。
	2. 当然 new 函数也可为引用类型分配内存，但这是不完整创建。以字典(map)为例，它仅仅分配了字典类型本身(实际上就是个指针包装)所需内存，
	   并没有分配键值存储内存，也没有初始化散列桶等内部属性，因此它无法正常工作。
*/

func main() {
	p := new(map[string]int)
	m := *p
	//m["a"] = 1 // 报错
	fmt.Println(m)
}

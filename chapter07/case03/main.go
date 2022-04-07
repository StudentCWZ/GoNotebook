/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 09:21
*/

package main

/*
	1. 可以像匿名字段那样，嵌入其他接口。
	2. 目标类型方法集中必须拥有包含嵌入接口方法在内的全部方法才算实现了该接口。
	3. 嵌入其他接口类型，相当于将其声明的方法集导入。这就要求不能有同名方法，因为不支持重载。
	4. 还有，不能嵌入自身或循环嵌入，那会导致递归错误。
*/

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct{}

func (*data) test() {}

func (data) string() string {
	return ""
}

func main() {
	var d data
	var t tester = &d
	t.test()
	println(t.string())
}

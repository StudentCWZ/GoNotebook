/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 08:45
*/

package main

/*
	1. 接口代表一种调用的契约，是多个方法声明的集合。
	2. 在某些动态语言里，接口也被称为协议(protocol)
	3. 接口除了解除依赖，有助于减少用户可视方法，屏蔽内部结构和实现细节。
	4. Go 接口实现机制很简洁，只要目标类型方法集内包含接口声明的全部方法，就被视为实现了该接口，无须做显式声明。
	5. 当然，目标类型可实现多个接口。
	6. 从内部实现来看，接口自身也是一种结构类型，只是编译器会对其做出很多限制。
       - 不能有字段
       - 不能定义自己的方法
       - 只能声明方法，不能实现
       - 可嵌入其他接口类型
	7. 接口通常以 er 作为名称后缀，方法名是声明组成部分，但参数名可不同或省略。
*/

type tester interface {
	test()
	string() string
}

type data struct{}

func (*data) test()         {}
func (data) string() string { return "" }

func main() {
	var d data
	//var t tester = d

	var t tester = &d
	t.test()
	println(t.string())
}






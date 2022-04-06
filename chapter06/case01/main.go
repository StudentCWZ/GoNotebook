/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 09:01
*/

package main

import "fmt"

/*
	1. 方法是与对象实例绑定的特殊函数。
	2. 方法是面向对象编程的基本概念，用于维护和展示对象的自身状态。
	3. 对象是内敛的，每个实例都有各自不同的独立特征，以属性和方法来暴露对外的通信接口。
	4. 普通函数则专注于算法流程，通过接收参数来完成特定的逻辑运算，并返回最终结果。
	5. 换句话说，方法有关联状态的，而函数通常没有。
	6. 方法和函数定义语法区别在于前者有前置实例接收参数(receiver)，编译器以此确定方法所属类型。
	7. 可以为当前包，以及除接口和指针以外的任何类型定义方法。
*/

type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}

func main() {
	var a N = 25
	println(a.toString())
}

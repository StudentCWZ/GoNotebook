/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 11:28
*/

package main

import (
	"fmt"
	"reflect"
)

/*
	1. 类型有一个与之相关的方法集(method set)，这决定了它是否实现某个接口。
		- 类型 T 方法集包含所有 receiver T 方法。
		- 类型 *T 方法集包含所有 receiver T + *T 方法
		- 匿名嵌入 S，T 方法集包含所有 receiver S 方法
		- 匿名嵌入 *S，T 方法集合包含所有 receiver S + *S 方法
		- 匿名嵌入 S 或 *S，*T 方法集包含所有 receiver S + *S 方法
	2. 方法集仅影响接口实现和方法表达式转换，与通过实例或实例指针调用的方法无关。
	3. 实例并不使用方法集，而是直接调用(或通过隐式字段名)。
	4. 面向对象的三大特征 "封装"、"继承"、"多态"，Go 仅实现了部分特征，它更倾向于 "组合优于继承" 这种思想。
	5. 将模块分解成相互独立的更小单元，分别处理不同方面的需求，最后以匿名嵌入方式组合到一起，共同实现对外接口。而且其简短一致的调用方式，
       更是隐藏了内部实现细节。
*/

type S struct{}

type T struct {
	S
}

func (S) sVal() {}

func (*S) sPtr() {}

func (T) tVal() {}

func (*T) tPtr() {}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var t T
	methodSet(t) // 显示 T 方法集
	println("-----------")
	methodSet(&t) // 显示 *T 方法集
}

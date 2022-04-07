/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 10:05
*/

package main

/*
	1. 接口使用一个名为 itab 的结构存储运行期所需的相关类型信息。
		type iface struct {
			tab *itab
			data unsafe.Pointer
		}

		type itab struct {
			inter *interfacetype // 接口类型
			_type *_type //实际对象类型
			fun [1]uintptr	// 实际对象方法地址
		}
	2. 利用调试器我们可查看这些结构存储的具体内容
*/

type Ner interface {
	a()
	b(int)
	c(string) string
}

type N int

func (N) a()               {}
func (*N) b(int)           {}
func (*N) c(string) string { return "" }

func main() {
	var n N
	var t Ner = &n

	t.a()
}

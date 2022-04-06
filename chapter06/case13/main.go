/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 14:59
*/

package main

/*
	1. 只要 receiver 参数类型正确，使用 nil 同样可以执行
*/

type N int

func (N) value() {

}
func (*N) pointer() {

}

func main() {
	var p *N

	p.pointer()         // method value
	(*N)(nil).pointer() // method value
	(*N).pointer(nil)   // method expression

	//p.value() // 报错: invalid memory address or nil pointer dereference
}

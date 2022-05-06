/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/6 10:28
*/

package main

import "fmt"

/*
	1. 变参本质上就是一个切片。只能接收一到多个同类型参数，且必须放在列表尾部。
	2. 将切片作为变参时，须进行展开操作。如果是数组，先将其转换成切片。
	3. 既然变参是切片，那么参数复制的仅是切片本身，并不包括底层数组，也因此可修改原数据。如果需要，可以内置函数 copy 复制底层数据。
*/

func testOne(s string, a ...int) {
	fmt.Printf("%T, %v\n", a, a)
}

func testTwo(a ...int) {
	fmt.Println(a)
}

func testThree(a ...int) {
	for i := range a {
		a[i] += 100
	}
}

func main() {
	testOne("abc", 1, 2, 3, 4)
	a := [3]int{10, 20, 30}
	testTwo(a[:]...)
	b := []int{10, 20, 30}
	testThree(b...)
	fmt.Println(b)
}

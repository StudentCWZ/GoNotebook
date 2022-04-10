/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/10 10:49
*/

package main

/*
	1. 所谓引用类型特指 slice、map、channel 这三种预定义类型。
	2. 相比数字、数组等类型，引用类型拥有更复杂的存储结构。
	3. 除分配内存外，它们还须初始化一系列属性，诸如指针、长度，甚至哈希分布、数据队列等。
	4. 内置函数 new 按指定类型长度分配零值内存、返回指针、并不关心类型内部构造和初始化方式。
	5. 而引用类型则必须使用 make 函数创建，编译器会将 make 转换为目标类型专用的创建函数(或指令)，以确保完成全部内存分配和相关属性初始化。
*/

func mkSlice() []int {
	s := make([]int, 0, 0)
	s = append(s, 100)
	return s
}

func mkMap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func main() {
	m := mkMap()
	println(m["a"])

	s := mkSlice()
	println(s[0])
}

package main

import "fmt"

/*
	1. 定义数组类型时，数组长度必须是非负整型常量表达式，长度是类型组成部分。也就是说元素类型相同，但长度不同的数组不属于同一个类型。
	2. 对于结构体复合类型，可省略元素初始化类型标签。
	3. 在定义多维数组时，仅第一维度允许使用 ... 。
	4. 内置函数 len 和 cap 都返回第一维度长度
	5. 如元素类型支持 ==、!= 操作符，那么数组也支持此操作
*/

func arrayType() {
	var d1 [3]int
	var d2 [2]int

	// d1 = d2	// 错误
	fmt.Printf("d1: %v d2: %v\n", d1, d2)
}

func arrayInit() {
	var a [4]int // 元素自动初始化为零

	b := [4]int{2, 5}         // 未提供初始值的元素自动初始化为 0
	c := [4]int{5, 3: 10}     // 可指定索引位置初始化
	d := [...]int{1, 2, 3}    // 编译器按初始值数量确定数组长度
	e := [...]int{10, 3: 100} // 支持索引初始化，但注意数组长度与此有关

	fmt.Printf("a: %v b: %v c: %v d: %v e: %v\n", a, b, c, d, e)
}

func structType() {
	type user struct {
		name string
		age  int8
	}
	d := [...]user{
		{"Tom", 20},
		{"Marry", 18},
	}
	fmt.Printf("%#v\n", d)
}

func multiArray() {
	a := [2][2]int{
		{1, 2},
		{3, 4},
	}

	b := [...][2]int{
		{10, 20},
		{30, 40},
	}

	c := [...][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{10, 20},
			{30, 40},
		},
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func multiArrayLenOrCap() {
	a := [2]int{0, 0}
	b := [...][2]int{
		{10, 20},
		{30, 40},
		{50, 60},
	}
	println(len(a), cap(a))
	println(len(b), cap(b))
	println(len(b[1]), cap(b[1]))
}

func ComparisonOperators() {
	var a, b [2]int
	println(a == b)
	c := [2]int{1, 2}
	d := [2]int{0, 1}
	println(c != d)

	var e, f [2]map[string]int
	// println(e == f)	// 错误：无法比较
	fmt.Println(e, f)
}

func main() {
	arrayType()
	arrayInit()
	structType()
	multiArray()
	multiArrayLenOrCap()
	ComparisonOperators()
}

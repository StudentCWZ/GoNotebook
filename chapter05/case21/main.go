package main

import "fmt"

/*
	1. 函数 len 返回当前键值对数量，cap 不接受字典类型。
	2. 因为内存访问安全和哈希算法等缘故，字典被设计成 not addressable，故不能直接修改 value 成员（结构体或者数组）
	3. 正确做法是返回整个 value，待修改后再设置字典键值，或直接用指针类型。
*/

func mapOperateIsFalse() {
	type user struct {
		name string
		age  byte
	}
	m := map[int]user{
		1: {"Tom", 19},
	}
	// m[1].age += 1	// 错误
	fmt.Println(m)
}

func mapOperateIsTrue() {
	type user struct {
		name string
		age  byte
	}
	m := map[int]user{
		1: {"Tom", 19},
	}
	u := m[1]
	u.age += 1
	m[1] = u // 设置整个 value
	fmt.Println("m:", m)

	m2 := map[int]*user{ // value 是指针类型
		1: {"Jack", 20},
	}
	m2[1].age++
	fmt.Println("m2:", m2) // m2[1] 返回的是指针，可透过指针修改目标对象
}

func main() {
	mapOperateIsFalse()
	mapOperateIsTrue()
}

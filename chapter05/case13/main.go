package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	1. 切片本身并非动态数组或数组指针。它内部通过指针引用底层数组，设定相关属性将数据读写操作限定在指定区域内
	2. 切片本身是个只读对象，其工作机制类似数组指针的一种包装。
	3. 可基于数组或数组指针创建切片，以开始和结束索引位置确定所引用的数组片段。不支持反向索引，实际范围是一个右半开区间。
	4. 属性 cap 表示切片所用数据切片片段的真实长度，len 用于限定可读的写元素数量。另外，数组必须 addressable，否则会引发错误
	5. 和数组一样，切片同样使用索引号访问元素内容。其实索引为 0，而非对应的底层数组真实索引位置。
	6. 可直接创建切片对象，无须预先准备数组。因为是引用类型，需使用 make 函数或显式初始化语句，它会自动完成数组内存分配。
*/

// type slice struct {
// 	array unsafe.Pointer
// 	len int
// 	cap int
// }

func sliceOperate() {
	m := map[string][2]int{
		"a": {1, 2},
	}
	// s := m["a"][:] 报错
	f := m["a"]
	s := f[:]
	fmt.Println(s)
}

func sliceIndex() {
	x := [...]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	s := x[2:5]
	for i := 0; i < len(s); i++ {
		println(s[i])
	}
}

func sliceMake() { // 注意：下面两种定义方式的区别：前者仅定义了一个 []int 类型变量，并未执行初始化操作，而后者则用初始化表达式完成了全部创建过程。
	s1 := make([]int, 3, 5)
	s2 := make([]int, 3)
	s3 := []int{10, 20, 5: 30}

	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s3, len(s3), cap(s3))
}

func sliceInit() {
	var a []int
	b := []int{}

	println(a == nil, b == nil)

	fmt.Printf("a: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a))) // 通过输出更详细的信息，我们可以看到两者的差异
	fmt.Printf("b: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
	fmt.Printf("a size: %d\n", unsafe.Sizeof(a))
}

func main() {
	sliceOperate()
	sliceIndex()
	sliceMake()
	sliceInit()
}

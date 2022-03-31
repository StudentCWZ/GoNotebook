package main

import "fmt"

/*
	1. 不支持比较操作，就算元素类型支持也不行，仅能判断是否为 nil
	2. 可获取地址，但不能像数组那样直接用指针访问元素内容
	3. 如果元素类型也是切片，那么就可以实现类似交错数组（jagged array）功能
*/

func testSlice() {
	a := make([]int, 1)
	b := make([]int, 1)
	// println(a == b) // 报错：无法比较
	fmt.Println(a, b)
}

func slicePtr() {
	s := []int{0, 1, 2, 3, 4}
	p := &s     // 取 header 地址
	p0 := &s[0] // 取 array[0] 地址
	p1 := &s[1]

	println(p, p0, p1)
	(*p)[0] += 100 // *[]int 不支持索引操作，须返回 []int 对象
	*p1 += 100     // 直接用元素指针操作
	fmt.Println(s)
}

func jaggedSlice() {
	x := [][]int{
		{1, 2},
		{10, 20, 30},
		{100},
	}
	fmt.Println(x[1])
	x[2] = append(x[2], 200, 300)
	fmt.Println(x[2])

}

func main() {
	testSlice()
	slicePtr()
	jaggedSlice()
}

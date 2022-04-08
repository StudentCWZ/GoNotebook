/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:46
*/

package main

import "fmt"

/*
	1. 切片可实现类似动态数组的功能。
*/

func main() {
	x := make([]int, 0, 5) // 创建容量为 5 的切片
	for i := 0; i < 8; i++ {
		x = append(x, i) // 追加数据。当超出容量限制时，自动分配更大的存储空间
	}
	fmt.Println(x)
}

package main

import "fmt"

/*
	1. 将切片视作 [cap]slice 数据源，据此创建新切片对象。不能超出 cap，但不受 len 限制
	2. 新建切片对象依旧指向原底层数组，也就是说修改对所有关联切片可见。
*/

func main() {
	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := d[3:7]
	s2 := s1[1:3]

	for i := range s2 {
		s2[i] += 100
	}
	fmt.Println(d)
	fmt.Println(s1)
	fmt.Println(s2)
}

package test

import "testing"

/*
	1. 在创建时预先准备足够空间有助于提升性能，减少扩张时的内存分配和重新哈希操作。
	2. 对于海量小对象，应直接用字典存储键值数据拷贝，而非指针。这有助于减少需要扫码的对象数量，大幅缩短垃圾回收时间。
	3. 另外，字典不会收缩内存，所以适当替换成新对象是必要的。
*/

func test() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = 1
	}
	return m
}

func testCap() map[int]int {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = 1
	}
	return m
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}

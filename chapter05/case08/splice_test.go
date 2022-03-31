package demo

import (
	"bytes"
	"strings"
	"testing"
)

/*
	1. 除类型转换外，动态构建字符串也容易造成性能问题。用加法操作符拼接字符串时，每次都须重新分配内存。如此，在构建超大字符串时，性能就显得极差。
	2. 改进思路是预分配足够的内存空间。常用方法是用 strings.Join 函数，它会统计所有参数长度，并一次性完成内存分配操作。
	3. 相对于加法操作符拼接字符串，strings.Join 方法提升了性能，另外 bytes.Buffer 也能完成类似操作，且性能想单。
	4. 对于数量较少的字符串格式拼接，可使用 fmt.Sprintf、text/template 等方法
*/

func test() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test()
	}
}

// testTwo 改进测试案例（strings.json）
func testTwo() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}

func BenchmarkTestTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = testTwo()
	}
}

// testThree 改进测试案例（bytes.Buffer）
func testThree() string {
	var b bytes.Buffer
	b.Grow(1000)

	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}

	return b.String()
}

func BenchmarkTestThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = testThree()
	}
}

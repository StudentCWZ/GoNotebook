package main

/*
	1. 字符串默认值不是 nil，而是 ""
	2. 使用 `` 定义不做转义处理的原始字符串（rwa string），支持跨行
	3. 注意：编译器不会解析原始字符串内的注释语句，且前置缩紧空格也属于字符串内容。
*/

func StringDefaultValue() {
	var s string
	println(s == "")
	// println(s == nil)	// 无效操作
}

func StringRawValue() {
	s := `line\r\n,
	line 2`
	println(s)
}

func main() {
	StringDefaultValue()
	StringRawValue()
}

/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/6 08:45
*/

package main

/*
	1. 函数是结构化编程的最小模块单元。它将复杂的算法过程分解为若干较小任务，隐藏相关细节，使得程序结构更加清晰，易于维护。
	2. 函数被设计成相对独立，通过接收输入参数完成一段算法指令，输出或存储相关结果。因此，函数还是代码复用和测试的基本单元。
	3. 关键字 func 用于定义函数。Go 中的函数有些不方便的限制，但也借鉴了动态语言的某些优点。
       - 无须前置声明
       - 不支持命令嵌套定义(nested)
       - 不支持同名函数重载(overload)
       - 不支持默认参数
       - 支持不定长变参
       - 支持多返回值
       - 支持命名返回值
       - 支持匿名函数和闭包
	4. 函数属于第一类对象，具备相同签名(参数及返回值列表)的视作同一类型。
	5. 从阅读和维护代码的角度来说，使用命名类型更加方便。
	6. 函数只能判断其是否为 nil ，不支持其他比较操作。
	7. 从函数返回局部变量指针是安全的，编译器会通过逃逸分析来决定是否在堆上分配内存。
	8. 在避免冲突的情况下，函数本着精简短小、望闻知意的原则。
       - 通常是动词和介词加上名词，例如 scanWords
       - 避免不必要的缩写，printError 要比 printErr 更好一些
       - 避免使用类型关键字，比如 buildUserStruct 看上去会很别扭
       - 避免歧义，不能有多种用途的解释造成的误解。
       - 避免只能通过大小写区分的同名函数
       - 避免与内置函数重名，这会导致误用
       - 避免使用数字，除非是特定专有名词，例如 UTF-8
       - 避免添加作用域提示前缀
       - 统一使用 camel/pascal case 拼写风格
       - 使用相同术语，保持一致性
       - 使用习惯用语，比如 init 表示初始化，is/has 返回布尔值结果
       - 使用反义词组命名行为相反的函数，比如 get/set、min/max 等
*/

func hello() {
	println("hello,world!")
}

func exec(f func()) {
	f()
}

func a() {}
func b() {}

func test() *int {
	a := 0x100
	return &a
}

func main() {
	f := hello
	exec(f)

	println(a == nil)
	//println(a == b)	// 报错

	var a *int = test()
	println(a, *a)
}

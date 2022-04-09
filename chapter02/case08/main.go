/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 14:59
*/

package main

/*
	1. 对变量、常量、函数、自定义类型进行命名，通常优先选用有实际含义，易于阅读和理解的字母或单词组合。
	2. 命名建议:
       - 以字母或下划线开始，由多个字母、数字和下划线组合而成。
       - 区分大小写
       - 使用驼峰拼写格式
       - 局部变量优先使用短名
       - 不要使用保留关键字
       - 不建议使用与预定义常量、类型、内置函数相同的名字
       - 专有名词通常会全部大写
	3. 尽管 Go 支持用汉字等 Unicode 字符命名，但从编程习惯上来说，这并不是好选择。
	4. 符号名字的首字母大小写决定了其作用域。首字母大些的为导出成员，可被包外引用，而小写则仅能在包内使用。
*/

func main() {
	var c int                 // c 代替 count
	for i := 0; i < 10; i++ { // i 代替 index
		c++
	}
	println(c)
}

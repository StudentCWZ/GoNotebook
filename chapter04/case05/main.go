/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/6 10:45
*/

package main

import (
	"errors"
	"fmt"
)

/*
	1. 有返回值的函数，必须有明确的 return 终止语句。
	2. 除非有 panic，或者无 break 的死循环，则无须 return 终止语句。
	3. 借鉴自动态语言的多返回值模式，函数得以返回更多状态，尤其是 error 模式。
	4. 稍有不便的是没有元组(tuple)类型，也不能用数组、切片接收，但可用 "_" 忽略掉不想要的返回值。多返回值可用作其他函数调用实参，或当作结果
       直接返回。
	5. 对返回值命名和简短变量定义一样，优缺点共存。
	6. 命名返回值让函数声名更加清晰，同时也会改善帮助文档和代码编辑器提示。
	7. 命名返回值和参数一样，可当作函数局部变量使用，最后由 return 隐式返回。
	8. 如果返回值类型能明确表明其含义，就尽量不要对其命名。
*/

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func divTwo(x, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("division by zero")
		return
	}
	z = x / y
	return // 相当于 "return z, err"
}

func logTest(x int, err error) {
	fmt.Println(x, err)
}

func testOne() (int, error) { // 多返回值用作 return 结果
	return div(5, 0)
}

func main() {
	logTest(testOne()) // 多返回值用作实参
}
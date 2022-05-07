/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/7 11:02
*/

package main

import (
	"fmt"
	"log"
)

/*
	1. 某些时候，我们需要自定义错误类型，以容纳更多上下文状态信息。这样的话，还可以基于类型做出判断。
	2. 自定义错误类型通常以 Error 为名称后缀。在用 switch 按类型匹配时，注意 case 顺序。应将自己定义类型放前面，优先匹配更具体的错误类型。
	3. 在正式代码中，我们不能忽略 error 返回值，应严格检查，否则可能会导致错误的逻辑状态。
	4. 调用多返回值函数时，除 error 外，其他返回值同样需要关注。
	5. 大量函数和方法返回 error，使得代码变得很难看，一堆堆检查语句充斥在代码行。解决思路有：
       - 使用专门的检查函数处理错误逻辑(比如记录日志)，简化检查代码。
       - 在不影响逻辑的情况下，使用 defer 延后处理错误错误状态(err 退化赋值)。
       - 在不中断逻辑的情况下，将错误作为内部状态保存，等最终"提交"时再处理。
*/

type DivError struct { // 自定义错误类型
	x, y int
}

func (DivError) Error() string { // 实现 error 接口方法
	return "division by zero"
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{x, y} // 返回自定义错误类型
	}
	return x / y, nil
}

func main() {
	z, err := div(5, 0)
	if err != nil {
		switch e := err.(type) { // 根据类型匹配
		case DivError:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(e)
		}
		log.Fatalln(err)
	}
	println(z)
}

/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/6 09:22
*/

package main

import (
	"fmt"
	"time"
)

/*
	1. Go 对参数的处理偏向保守，不支持有默认值的可选参数，不支持命名实参。调用时，必须按签名顺序传递指定类型和数量的实参，就算以 "_" 命名参数
       也不能忽略。
    2. 在参数列表中，相邻的同类型参数可合并。
	3. 参数可视作函数局部变量，因此不能在相同层次定义同名变量。
	4. 形参是指函数定义中的参数，实参则是函数调用时所传递的参数。形参类似函数局部变量，而实参则是函数外的对象，可以是常量、变量、表达式或函数
       等。
	5. 不管是指针、引用类型，还是其他类型参数，都是值拷贝传递。区别无非就是拷贝目标对象，还是拷贝指针而已。
	6. 函数调用前，会为形参和返回值分配内存空间，并将实参拷贝到形参内存。
	7. 表面上看，指针参数的性能要更好一些，但实际上得具体分析。被复制的指针会延长目标对象生命周期，还会导致它分配到堆上，那么其性能消耗就得加上
       堆内存分配和垃圾回收的成本。
	8. 其实在栈上复制小对象只须很少的指令即可完成，远比运行时进行堆内存分配要快得多。另外，并发编程也提倡尽可能使用不变对象(只读或复制)，这可
       消除数据同步等麻烦。当然，如果复制成本很高，或需要修改原对象状态，自然使用指针更好。
*/

func test(x, y int, s string, _ bool) *int {
	return nil
}

func testTwo(x *int) {
	fmt.Printf("pointer: %p, target: %v\n", &x, x)
}

func testThree(p *int) {
	go func() { // 延长 p 生命周期
		println(p)
	}()
}

func main() {
	test(1, 2, "abc", false)
	a := 0x100
	p := &a
	fmt.Printf("pointer: %p, target: %v\n", &p, p) // 从输出结果可以看出，尽管实参和形参指向同一目标，传递指针时依然被复制
	testTwo(p)

	x := 100
	p1 := &x
	testThree(p1)
	time.Sleep(time.Second * 2)
}

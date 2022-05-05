/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/5 10:50
*/

package main

/*
	1. switch 与 if 类似，switch 语句也用于选择执行，但具体使用场景会有所不同。
	2. 条件表达式支持非常量值，这要比 C 更加灵活。相比 if 表达式，switch 值列表要更加简洁。
	3. 编译器对 if、switch 生成的机器指令可能完全相同，所谓谁性能更好须看具体情况，不能作为主观判断条件。
	4. switch 同样支持初始化语句，按从上到下、从左到右顺序匹配 case 执行。只有全部匹配失败时，才会执行 default 块。
	5. 相邻的空 case 不构成多条件匹配。
	6. 无须显式执行 break 语句，case 执行完毕后自动中断。
	7. 如须贯通后续 case，须执行 fallthrough，但不再匹配后续条件表达式。
	8. 注意，fallthrough 必须放在 case 块结尾，可使用 break 语句阻止。
	9. 某些时候，switch 还被用来替换 if 语句。被省略的 switch 条件表达式默认为 true，继而与 case 比较表达式结果匹配。
	10. switch 语句也可用于接口类型匹配。
*/

func switchCaseOne() {
	a, b, c, x := 1, 2, 3, 2

	switch x {
	// 将 x 与 case 条件匹配
	case a, b:
		println("a | b") // 多个匹配条件命中其一即可(or)，变量
	case c:
		println("c") // 单个匹配条件
	case 4:
		println("d") // 常量
	default:
		println("z")
	}
}

func switchCaseTwo() {
	switch x := 5; x {
	default:
		x += 100
		println(x)
	case 5:
		x += 50
		println(x)
	}
}

func switchCaseThree() {
	x := 10
	switch x {
	case 8:
	// 单条件，内容为空隐式 "case8: break;"
	case 10:
		println(x)

	}
}

func switchCaseFour() {
	switch x := 5; x {
	default:
		println(x)
	case 5:
		x += 10
		println(x)

		fallthrough // 继续执行下一 case，但不再匹配条件表达式
	case 6:
		x += 20
		println(x)

		//fallthrough // 如果在此继续 fallthrough，不会执行 default，完全按源码顺序；导致 "cannot fallthrough final case in switch" 错误
	}
}

func switchCaseFive() {
	switch x := 5; x {
	case 5:
		x += 10
		println(x)

		if x >= 15 {
			break
		}
		fallthrough
	case 6:
		x += 20
		println(x)
	}
}

func switchCaseSix() {
	switch x := 5; {
	case x > 5:
		println("a")
	case x > 0 && x <= 5:
		println("b")
	default:
		println("z")
	}
}

func main() {
	switchCaseOne()
	switchCaseTwo()
	switchCaseThree()
	switchCaseFour()
	switchCaseFive()
	switchCaseSix()
}

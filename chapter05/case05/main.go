package main

import "fmt"

func StringCircle() {
	s := "雨痕"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%d: [%c]\n", i, s[i])
	}
	for i, c := range s { // rune：返回数组索引号，以及 Unicode 字符
		fmt.Printf("%d: [%c]\n", i, c)
	}
}

func main() {
	StringCircle()
}

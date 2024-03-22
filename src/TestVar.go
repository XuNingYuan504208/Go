package main

import (
	"fmt"
)

var staticA int = 100

var staticB = 200

// := 初始化方式只能在函数中使用
// staticC := 500;

func declare() {
	var a int = 1
	var b = 1

	c := 1

	fmt.Println("a:", a, "b:", b, "c:", c)
	fmt.Println("staticA:", staticA, "staticB:", staticB, "c:", c)

	// 多变量声明
}

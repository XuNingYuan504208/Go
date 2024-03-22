package main

import (
	"fmt"
)

func foo1(a string, b int) int {
	fmt.Println("a : ", a, "b : ", b)
	c := b * 2
	return c
}

// 返回多个返回值 ，匿名的
func foo2(a string, b int) (int, int) {

	return 666, 777
}

// 返回多个返回值，有形参名称
// func foo4(a string, b int) (r1, r2 int) {} 省略写法
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println(a, b)
	r2 = 100
	var r3 = 200
	println(r3)
	return r2, r3
}

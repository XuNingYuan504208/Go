package Goroutine

import (
	"runtime"
	"time"
)

func main() {
	println("main routine start ...")

	// 创建新的 goroutine，使用匿名函数
	go func() {
		defer println("anonymous function A defer end ...")
		//return

		// 第二个goroutine
		go func() {
			defer println("anonymous function B defer end ...")
			// 只是退出当前子函数
			//return
			// 退出当前goroutine

			/*
				goroutine anonymous func A start ...
				anonymous function A defer end ...
				anonymous function B defer end ...
			*/
			runtime.Goexit()

			println("goroutine anonymous func B start ... ")
		}()
		println("goroutine anonymous func A start ... ")
	}()

	// 死循环
	for {
		//
		time.Sleep(1 * time.Second)
	}

}

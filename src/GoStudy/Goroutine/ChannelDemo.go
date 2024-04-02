package Goroutine

import "time"

func main() {
	ch := make(chan int, 3)

	go func() {
		defer println("goroutine defer end")
		/*
			i > ch 的capacity时，会阻塞住，直到main 从ch里拿走元素
			打印结果：
			goroutine running : len(c) :  1
			goroutine running : len(c) :  2
			goroutine running : len(c) :  3
			main :: i :  0 ; chVal :  5
			goroutine running : len(c) :  3
			goroutine defer end
			main :: i :  1 ; chVal :  6
			main :: i :  2 ; chVal :  7
			main end

		*/
		for i := 0; i < 4; i++ {
			// i < 3的时候goroutine不会阻塞
			//for i := 0; i < 3; i++ {
			ch <- i + 5
			println("goroutine running : len(c) : ", len(ch))
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		chVal := <-ch
		println("main :: i : ", i, "; chVal : ", chVal)
	}
	println("main end")
	// main hold
	for {
		time.Sleep(1 * time.Second)
	}
}

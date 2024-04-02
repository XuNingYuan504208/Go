package Goroutine

import "time"

func ChannelWithCapacityDemo() {
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

/*
channel value :  5
channel value :  6
channel value :  7
channel value :  8
channel value :  9
channel closing ...
main ending ...
main end
*/
func CloseChannelDemo() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i + 5
		}
		time.Sleep(1 * time.Second)
		println("channel closing ...")

		// close关闭一个channel
		// 如果不关闭，则main会一直等待关闭，会打印出channel全部数据，但是由于goroutine没有继续写所以main会阻塞住
		// 报错：all goroutines are asleep - deadlock 死锁
		close(ch)
	}()

	for {
		// ok 如果为true：channel没有关闭；如果false：channel已经关闭
		// 先执行前面表达式，最后if判断的依据是ok
		if value, ok := <-ch; ok {
			println("channel value : ", value)
		} else {
			break
		}
	}
	// 上面for可以改成
	// range 会阻塞住，一直拿channel的数据
	for data := range ch {
		println(data)
	}
	println("main ending ...")
	time.Sleep(1 * time.Second)
	println("main end ")
}

func SelectChannelDemo() {
	fi := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			println(<-fi)
		}
		// 向quit 写一个0
		quit <- 0
	}()

	selectDemo(fi, quit)
}

func selectDemo(fi, quit chan int) {
	x := 0
	for {
		select {
		case fi <- x:
			{
				println("write x + 5")
				x += 5
			}
		case <-quit:
			println("case quit ")
			return
		default:

		}
	}
}

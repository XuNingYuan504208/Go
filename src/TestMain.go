package main

func mainTest() {
	//declare()
	//fmt.Println("BEIJING:", BEIJING, "SHANGHAI:", SHANGHAI, "SHENZHEN:", SHENZHEN)
	//c := foo1("aa", 5)

	/*var ret1 int = 0
	ret2 := 5
	println(ret1, ret2)
	ret1, ret2 = foo3("aa", 5)
	println(ret1, ret2)
	//println(c)

	var str string

	println("str : ", str)*/

	//lib1.Lib1Test()
	//lib2.Lib2Test()
	//myLib2.Lib2Test()

	//Lib2Test()
	//a := 1
	//println("a : ", a)
	//println("&a : ", &a)

	// & 代表传a 的地址引用
	//changeV(&a)
	//println(a)
	//var a, b = 1, 2
	//swap(&a, &b)
	//println("a:", a, "b:", b)

	//println("11111")
	//println("22222")
	//println("33333")
	//defer println("end ...")
	//println("33333")

	/*println("main routine start ...")

	// 创建新的 goroutine，使用匿名函数
	go func() {
		defer println("anonymous function A defer end ...")
		//return

		// 第二个goroutine
		go func() {
			defer println("anonymous function B defer end ...")
			// 只是退出当前子函数
			//return
			// 退出当前goroutine，结果：打印B defer 和A defer
			runtime.Goexit()
			println("goroutine anonymous func B start ... ")
		}()
		println("goroutine anonymous func A start ... ")
	}()

	// 死循环
	for {
		//
		time.Sleep(1 * time.Second)
	}*/

	/*ch := make(chan int, 3)

	go func() {
		defer println("goroutine defer end")

		for i := 0; i < 4; i++ {
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

	for {
		time.Sleep(1 * time.Second)
	}*/

}

/*
*
changeV的参数为指针，只接受传地址
*/
func changeV(p *int) {
	println("*p : ", *p)
	println("p : ", p)
	*p = 10
}

/*
交换 a,b
*/
func swap(pa *int, pb *int) {
	println("exchange :: before pa:", pa, ",b:", pb)
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
	println("exchange :: pa:", pa, ",pb:", pb)

}

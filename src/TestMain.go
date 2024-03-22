package main

func main() {
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

	println("11111")
	println("22222")
	println("33333")
	defer println("end ...")
	println("33333")
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

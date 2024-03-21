package GoStudy

import (
	"fmt"
	"time"
)

func goFunc(i int) {

	fmt.Println("i:", i)
}

func main() {
	for i := 0; i < 1000; i++ {
		go goFunc(i) // go 关键字开启协程
	}
	time.Sleep(time.Second)
}

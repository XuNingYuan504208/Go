package goStruct

// 接口
type Animal interface {
	Sleep()
	GetColor()
}

// 具体实现 1
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	println("Cat sleeping ...")
}

func (this *Cat) GetColor() {
	println("Cat get color ...")
}

// 具体实现  2
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	println("Dog sleeping ...")
}
func (this *Dog) GetColor() {
	println("Dog get color ...")
}

func main() {
	var animal Animal
	animal = &Cat{"yellow"}
	animal.Sleep()
	animal.GetColor()
}

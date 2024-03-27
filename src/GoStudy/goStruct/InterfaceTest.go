package goStruct

// 接口
type Animal interface {
	Sleep()
	GetColor()
}

// 具体实现
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	println("Cat sleeping ...")
}

func (this *Cat) GetColor() {
	println("Cat get color ...")
}

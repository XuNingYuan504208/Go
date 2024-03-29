package goInterface

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

type Person struct {
}

func (p *Person) ReadBook() {
	println("person read book ...")
}
func (p *Person) WriteBook() {
	println("person write book ...")
}

func main() {
	// p {type : Person ,value : Person{}地址 }
	p := &Person{}
	// r {type:nil , value : nil}
	var r Reader
	// r {type : Person ,value : Person{}地址 }
	r = p

	r.ReadBook()

	var w Writer
	// 断言成功：因为w r 具体type是一致的
	w = r.(Writer)

	w.WriteBook()

	_, ok := r.(Writer)
	if ok {
		println("assertion success")
	}
}

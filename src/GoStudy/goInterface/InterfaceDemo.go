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
}

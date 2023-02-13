package main

type Foo struct {
	v int
}

func NewFoo(n *int) Foo {
	print(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	print(*n)
}

func main03() {
	var x = 1

	var p = &x

	//fmt.Println(*p)

	defer NewFoo(p).Bar(p)
	x = 2
	p = new(int)

	//fmt.Println("----: ", *p)
	NewFoo(p)
}

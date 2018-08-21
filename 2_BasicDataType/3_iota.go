package main

const (
	a = iota
	b
	c
)

const (
	_ = iota
	d
)

const (
	e = 1 << iota
	f = 2 << iota  //01->move two bit to the left->100(is 4)
	g
)

func main() {

	println("a:",a,"b:",b,"c:",c)  //0,1,2
	
	println("d:",d)  //1
	
	println("e:",e,"f:",f,"g:",g)  //1,4,8
}

package main

var a = 1  //global variable

func main() {

	println("main()-a:",a)
	eg1()  //1
	
	a1 := 2
	println("main()-a1:",a1)
	//eg2()  //undefined
	
	//println("local-a4:",a4)  //undefined
	println("global-a3:",a3)
	a4 := 3
	println("local-a4:",a4)
	
	
}
func eg1(){
	println("eg1()-a:",a)
}

func eg2() {
	//println("eg2()-a1:",a1)
}

var a3 = 3

package main

import (
	"fmt"
)

func main() {
	
	a,b,c := 'a',1,1.1
	
	cc := 222.2
	ccc := 0.0666
	
	d,e,f := true,"hello go!",`how are you?`
	
	var g int
	var h string

	println("a:",a,"b:",b,"c:",c,"d:",d,"e:",e,"f:",f,"g:",g,"h:",h)  //a:97,g:0,h:""
	println("c",c);		//+1.100000e+000
	println("cc",cc);	//+2.222000e+002
	println("ccc",ccc);	//+6.660000e-002
	//data type
	fmt.Printf("%T \n", a)	//int,(not string)
	fmt.Printf("%T \n", b)	//int32
	fmt.Printf("%T \n", c)	//float64
	fmt.Printf("%T \n", d)	//bool
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	
	var student1 []string
	student2 := []string{}
	
	fmt.Println(student1 == nil)  //true
	fmt.Println(student2 == nil)  //false
}

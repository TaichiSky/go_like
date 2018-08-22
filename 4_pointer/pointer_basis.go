package main

import (
	f "fmt"
)

const PI float64 = 3.14

func main() {

	a := 10
	f.Println("the value of a is:",a,",and the &a is :",&a)
	b := &a
	f.Println("the value of b is:",*b,",and the b is :",b,",and the &b is :",&b)
	
	f.Println("update the value of b")
	*b = 9
	f.Println("the value of a is:",a,",and the &a is :",&a)
	f.Println("the value of b is:",*b,",and the b is :",b,",and the &b is :",&b)
	
	c := 2
	changeValue1(2)
	f.Println("exec changeValue1...,c is :",c)
	changeValue2(&c)
	f.Println("exec changeValue2...,c is :",c)
	
	var r float64
	f.Print("Enter the radius of circle: ")
	f.Scan(&r)
	area := r*r*PI
	f.Println("Area of the circle is :",area)
}
func changeValue1(value int) {
	value = 3
}
func changeValue2(value *int) {
	*value = 3
}

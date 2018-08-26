package main

import (
	f "fmt"
)

var b int

func main() {
	a:= 1
	f.Println("method-var_method-print a:",a)
	{
		f.Println("method-var_local-print a:",a)
	}
	
	f.Println("call method add1() firstly:",add1())
	f.Println("call method add1() secondly:",add1())
	
	c := 0
	add2 := func() int {
		c++
		return c
	}
	f.Println("call method add2() firstly:",add2())
	f.Println("call method add2() secondly:",add2())
	
	add3 := addWrapper()
	add33 := addWrapper()
	f.Println("call method add3() firstly:",add3())  //1
	f.Println("call method add3() secondly:",add3())  //2
	f.Println("call method add33() firstly:",add33())  //1
	f.Println("call method add33() secondly:",add33())  //2
	
	//callback test
	e := 2
	call_test1(e, func(num1 int) {
		f.Println("e:",num1)
	})
	
	//filter of callback test
	g := []int{1,2,3}
	arrays := call_test2(g, func(item1 int) bool {
		return item1 > 1
	})
	f.Println("After filter f:",arrays)
	
	//recursion test
	h := 5
	f.Println("recursion 5 is :",fac(h))
}

func add1() int {
	b++
	return b
}

func addWrapper() func() int {
	d := 0
	return func() int {
		d++
		return d
	}
}

func call_test1(num int, callback func(int)) {
	num += 3
	callback(num)
}

func call_test2(array []int, callback func(int) bool) []int {
	var arr []int
	for _, item := range array {
		if callback(item) {
			arr = append(arr, item)
		}
	}
	return arr
}

func fac(num int) int {
	if num == 0 {
		return 1
	}
	return num * fac(num-1)
}

package main

import (
	f "fmt"
)

func main() {
	a := []int{1,2,3}
	f.Println("a:",a,"a.length:",len(a))
	
	//capacity is assigned
	a_c1 := []int{1}
	a_c2 := []int{1,2}
	a_c3 := []int{1,2,3}
	a_c4 := []int{1,2,3,4}
	a_c5 := []int{1,2,3,4,5}
	f.Println("cap of a_c1:",cap(a_c1))
	f.Println("cap of a_c2:",cap(a_c2))
	f.Println("cap of a_c3:",cap(a_c3))
	f.Println("cap of a_c4:",cap(a_c4))
	f.Println("cap of a_c5:",cap(a_c5))
	
	//the array capacity is double that of when it is auto extension
	for i := 2; i<6; i++ {
		a_c1 = append(a_c1, i)
		f.Println("cap of a_c1:",cap(a_c1))
	}
	
	b := make([]int, 2, 3)	//len=2,cap=3
	f.Println("b:",b,",b.length:",len(b),",cap of b:",cap(b))
	
	b_plus := make([]int, 3)  //len=cap=3
	f.Println("b_plus:",b_plus,",b_plus.length:",len(b_plus),",cap of b_plus:",cap(b_plus))
	
	//b[2] = 1  //index out of range
	b = append(b, 1,2,3)  //auto extension cap
	for index,item := range b {
		f.Println("b[",index,"] = ",item)
	}
	f.Println("b:",b,",b.length:",len(b),",cap of b:",cap(b))
	
	//the max of slice_index must lower to cap(array)
	f.Println("slice b [index:index-1]:",b[2:4])
	f.Println("slice b [index]:",b[2])
	f.Println("slice b [index:]:",b[2:])
	f.Println("slice b [:index-1]:",b[:4])
	b_copy := b[:]
	f.Println("slice b [:]:",b_copy)
	
	//append slice
	c := []int{4,5,6}
	b = append(b,c...)
	f.Println("slice c add to b:",b)
	
	//delete
	b = append(b[2:5], c[1:3]...)
	f.Println("slice b,delete 0,0,4:",b)
}

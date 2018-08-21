package main

import (
	"fmt"
)

func main() {
	
	var a int
	b := 3
	c := [5]int{1,3,5,7}
	
	print("---循环[0-5)的a值：")
	for a := 0; a<5; a++ {
		print(a,"\t")
	}
	
	print("\n---若a<b=3就自动加1并输出：")
	for a <b {
		a++
		print(a,"\t")
	}
	
	println("\n---range输出数组c的值：")
	for i,x := range c {
		println("索引下标",i,"的值是：",x)
	}
	
	d,e,f := 3,2,2
	
	println("---continue test")
	for d < 7 {
		if d == 4 {
			d += 1
			continue
		}
		println("d的值是：",d)
		d++
	}
	
	println("---break test")
	for e < 6 {
		if e == 5 {
			break
		}
		println("b的值是：",e)
		e++
	}
	
	println("---print f with no 'for' condition")
	for {
		print(f,"\t")
		if f > 4 {
			break
		}
		f++
	}
	
	if g := 1 ; true {
		println("\nbool condition's g:",g)
	}
	
  //arrays sort
	h := [6]int{96,4,54,23,6,18}
	fmt.Println("before order:",h)
	for i,len := 0,len(h); i<len-1; i++ {
		for j := i+1;j<len; j++ {
			if h[i] > h[j] {
				h[i],h[j] = h[j],h[i]
			}
		}
	}
	fmt.Println("after order:",h)
}

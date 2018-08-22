package main

import (
	f "fmt"
	"bytes"
)

func main() {
	sayHello("hello","guy")

	f.Println(sayHello1("hello","guy"))
	
	f.Println(sayHello2("hello","guy"))
	
	f.Println(sayHello3("zhang","li","zhao","liu"))
	strArray := []string{"zhang","li","zhao","liu"}
	f.Println(sayHello3(strArray...))
	f.Println(sayHello33(strArray))
}
//func
func sayHello(say, name string) {
	f.Println(say,",",name)
}
//params-type
func sayHello1(say, name string) string {
	//return f.Println(say,",",name)  //too many args to return
	return f.Sprint(say,",",name)
}
//naming-return
func sayHello2(say, name string) (returnStr string) {
	returnStr = f.Sprint(say,",",name)
	return
}
//multi-params
func sayHello3(strs ...string) string {
	var namesBuffer bytes.Buffer
	for _, item := range strs {
		namesBuffer.WriteString(item)
	}
	return namesBuffer.String()
}
func sayHello33(strs []string) string {
	var namesBuffer bytes.Buffer
	for _, item := range strs {
		namesBuffer.WriteString(item)
	}
	return namesBuffer.String()
}

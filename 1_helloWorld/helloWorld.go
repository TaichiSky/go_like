package main

func main() {
	greet1 := "Hello,go!"
	println("I saying:",greet1)
	
	greet2 := takeGreets()
	println(greet2())
}
func takeGreets() func() string {
	return func() string {
		return "Hello go!"
	}
}

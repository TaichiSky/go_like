package main

func main() {

	println("---switch test")
	switch 1 {
	case 0:
		println("You are 0")
	case 1:
		println("You are 1")
	case 2:
		println("You are 2")
	default:
		println("Who are you?")
	}
	
	/**
	 * fallthrough在匹配的条件中使用，这里匹配的是1,case 2语句也会打印
	 * 必须从匹配的条件开始且连续使用n个fallthrough才会依次打印n+1个语句
	 */
	println("---fallthrough of switch test")
	switch 1 {
	case 0:
		println("You are 0")
		fallthrough	//invalid
	case 1:
		println("You are 1")
		fallthrough
	case 2:
		println("You are 2")
	case 3:
		println("You are 3")
		fallthrough	//invalid
	default:
		println("Who are you?")
	}
	
	println("---multi_match of switch test")
	switch 3 {
	case 0:
		println("You are 0")
	case 1:
		println("You are 1")
	case 2,3:
		println("You are 2")
	default:
		println("Who are you?")
	}
	
	println("---no value of switch test")
	a := 0
	switch {
	case a == 0:
		println("You are 0")
	case a == 1:
		println("You are 1")
	case a == 2:
		println("You are 2")
	default:
		println("Who are you?")
	}
}

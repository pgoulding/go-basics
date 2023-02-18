package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println(("7 is odd"))
	}

	if 8%4 == 0 { // you can have an if statement without an else
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // a statement can procede conditionals; any variables declared in this styatement are available in the current and all subsequent branches
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

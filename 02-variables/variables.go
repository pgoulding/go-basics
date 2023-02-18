package main

import "fmt"

func main() {
	var a = "initial" // var delares 1 or more variables
	fmt.Println(a)

	var b, c int = 1, 2 // You can defclare multiple variables at once
	fmt.Println(b, c)
	var d = true // Go will infer the type of initialized variables
	fmt.Println(d)

	var e int // Variables declared without a corresponding intialization are zero-valued. For example the zero-value for an int is 0
	fmt.Println(e)

	f := "apple" // the := syntax is short hand for declaring and initializing a variable. This syntax is only available insde functions
	fmt.Println(f)
}

package main

import "fmt"

func main() {
	fmt.Println("for")

	i := 1
	for i <= 3 { // the most basic type of for loop with one condition
		fmt.Println("i: ", i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { // a classic initial/condition/after for loop
		fmt.Println("j: ", j)
	}

	for { // for without conditon will loop repeatedly until you break out of the loop or return from the enclosing function
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // you can also continue to the next iteration fo the loop
		}
		fmt.Println("n: ", n)
	}
}

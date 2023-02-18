package main

import (
	"fmt"
	"math"
)

const s string = "constant" // const declares a constant variable

func main() {
	fmt.Println(s)

	const n = 500000000 // a const statement can appear anywhere a variable can

	const d = 3e20 / n // constant expressions perform arithmetic with arbitrary precision
	fmt.Println(d)

	fmt.Println(int64(d)) // a numeric constant has no type until its given one, such as by an explicit conversion

	fmt.Println(math.Sin(n)) // a number can be given a type by usit it in a context that requires one, such as a varibale assignment or function call. For example, here math.Sin expects a float64
}

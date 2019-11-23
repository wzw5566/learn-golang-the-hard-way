package main

import "fmt"

func main() {
	var num = 5.5 // Type inferred as `float64` (the default type for floating-point numbers)

	fmt.Printf("num type is %T\n",num)
}

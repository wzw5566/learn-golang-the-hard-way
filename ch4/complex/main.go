package main

import "fmt"

func main() {
	var x = 5 + 8i // Type inferred as `complex128`
	var a complex128 = complex(6, 2) // Create the complex number use complex function
	var b complex64 = complex(9, 2)


	fmt.Printf("The type of x is %T\n", x)
	fmt.Printf("The type of a is %T\n", a)
	fmt.Printf("The type of b is %T\n", b)
}

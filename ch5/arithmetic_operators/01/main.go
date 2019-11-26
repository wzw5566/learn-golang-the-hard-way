package main

import "fmt"

func main() {
	var A int = 10
	var B int = 20
	var C int

	C = A + B
	fmt.Println("A + B =", C)

	C = A - B
	fmt.Println("A - B =", C)

	C = A * B
	fmt.Println("A * B =", C)

	C = A / B
	fmt.Println("A / B =", C)

	C = A % B
	fmt.Println("A % B =", C)

	A++
	fmt.Println("A++ =", A)

	A--
	fmt.Println("A-- =", A)

}

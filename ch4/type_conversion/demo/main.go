package main

import "fmt"

func main() {

	var a int64 = 4
	var b int = int(a) // Explicit Type Conversion

	var c float64 = 6.5

	// Explicit Type Conversion
	var result = float64(b) + c // Works

	fmt.Println("result is", result)

}

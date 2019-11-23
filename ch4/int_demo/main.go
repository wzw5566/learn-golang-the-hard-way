package main

import "fmt"

func main() {
	// When you don't declare any type explicitly, the type inferred is `int`
	// The default type for integers
	var myInt8 int8 = 100

	var myInt int = 1200

	var myUint uint = 400

	var myHexNumber = 0xFF // Use prefix '0x' or '0X' for declaring hexadecimal numbers

	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	fmt.Printf("%d, %d, %d, %#x, %#o\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber)

	fmt.Printf("%T, %T, %T, %T, %T\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber)

}

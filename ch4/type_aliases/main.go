package main

import "fmt"

func main() {

	// Type inferred as `rune` (Default type for character values)
	// Alias For Int32
	var theV = 'V'

	// Alias For uint8
	var c byte = 255  

	fmt.Printf("theV type is %T, c type is %T \n", theV,c)
}

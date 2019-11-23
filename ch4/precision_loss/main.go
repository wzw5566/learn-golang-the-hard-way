package main

import "fmt"

func main() {
	var num1 float32 = -123.0000901 // Loss of precision 
	var num2 float64 = -123.0000901 

	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

}

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Println("num1=", num1, "num2=", num2)
	fmt.Println("num1 size is", unsafe.Sizeof(num1), "num2 size is", unsafe.Sizeof(num2))

}

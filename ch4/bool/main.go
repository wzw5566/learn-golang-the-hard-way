package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str1 := "LearnGo"
	str2 := "learnGo"
	str3 := "LearnGo"
	result1 := str1 == str2 // bool size is 1 byte
	result2 := str1 == str3

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("The type of result1 is %T and "+
		"the type of result2 is %T\n",
		result1, result2)
	fmt.Println("result1 size =", unsafe.Sizeof(result1))
}

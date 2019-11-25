package main

import "fmt"

func main() {
	var i int = 20
	fmt.Println("i address =", &i)

	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr value=%v\n", *ptr)
	fmt.Printf("ptr address= %v\n", &ptr)

}

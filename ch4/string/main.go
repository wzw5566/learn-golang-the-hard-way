package main

import "fmt"

func main() {

	title := "\tlearn golang the hard way!"

	str1 :=
		`
		package main
		import "fmt"
		func main() {
		var num1 float32 = -123.0000901 // Loss of precision 
		var num2 float64 = -123.0000901 
		fmt.Println("num1=", num1)
		fmt.Println("num2=", num2)
		}
	`
	fmt.Println(title)
	fmt.Printf(str1)
}

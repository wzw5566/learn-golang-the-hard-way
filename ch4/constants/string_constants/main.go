package main

import "fmt"

func main() {
	var defaultName = "Vincent" //allowed
	type myString string
	var customName myString = "Vincent" //allowed,String Constants
	//customName = defaultName        //not allowed

	fmt.Println(defaultName,customName)

}

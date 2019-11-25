package main

import "fmt"

func main() {
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed
	//defaultBool = customBool          //not allowed

	fmt.Println(defaultBool, customBool)
}



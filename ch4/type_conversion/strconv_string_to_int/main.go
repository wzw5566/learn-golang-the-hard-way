package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s string = "1234"
	i, _ := strconv.Atoi(s) // To convert from string to int

	fmt.Println(i)                 // 1234
	fmt.Println(reflect.TypeOf(i)) // int
}

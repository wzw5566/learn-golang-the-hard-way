package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var j int = 1234
	k := strconv.Itoa(j) // To convert from int to string

	fmt.Println(k)                 //1234
	fmt.Println(reflect.TypeOf(k)) //string
}

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i int64 = 1234
	s := strconv.FormatInt(i, 10)  // To convert from int64 to string

	fmt.Println(s)                 //1234
	fmt.Println(reflect.TypeOf(s)) //string
}

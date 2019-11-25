package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s string = "123456"
	i, err := strconv.ParseInt(s, 10, 64)  // To convert from string to int64
	if err != nil {
		panic(err)
	}

	fmt.Println(i) // 123456
	fmt.Println(reflect.TypeOf(i)) //int64
}


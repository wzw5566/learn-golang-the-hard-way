package main

import "fmt"

// Weekday, iota :Elegant
const (
	Sunday = iota +1   // 1
	Monday             // 2
	Tuesday            // 3
	Wednesday          // 4
	Thursday           // 5
	Friday             // 6
	Saturday           // 7
)

func main() {
	fmt.Println(Friday)

}

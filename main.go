package main

import "fmt"

// Integer ...
type Integer int32

func checkIf(i Integer) bool {
	return i > 10
}

func main() {
	fmt.Println("Hello World")
	hello()
	fmt.Println(checkIf(11))

	var num1 Integer = 123
	var ptr = &num1

	fmt.Printf("%x=%d", ptr, *ptr)
}

func hello() {
	fmt.Println("Hello World 2")
}

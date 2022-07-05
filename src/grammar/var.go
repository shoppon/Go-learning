package main

import "fmt"

func main() {
	var s string
	s = "hello"
	var _integer int
	_integer = 10
	fmt.Println(s)
	fmt.Println(_integer)

	_integer1 := 11
	fmt.Println(_integer1)

	_pointer := &_integer1
	fmt.Println(_pointer)
}

package main

import "fmt"

func main() {
	var value int8 = 123
	var new_value int16 = int16(value)
	fmt.Println(value)
	fmt.Println(new_value)
}

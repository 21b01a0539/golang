package main

import "fmt"

//Gererics are used to work with two or kore datatypes
func add[T int|float64](x T, y T) T {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.5, 2.5))
	fmt.Println(add(2.9, 3.7))	
}
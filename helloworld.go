package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	//varibles declaration
	var a int = 10
	var name string  //declaring with datatype
	name = "John"
	fmt.Println(a)
	fmt.Println(name)
	var x int
	fmt.Println(x)
	y := "svecw"
	a, b, _, f := 1, 2, 3, "bhargavi"
	fmt.Println(a, b, f)
	fmt.Println(y)
	
	//Binary hexadecimal octal
	fmt.Printf("%b %x %o", 42, 42, 42)
	//%b is binary, %x is hex and %o is octal
	
	//variables and type conversion
	var b1 int = 42
	var b2 int32
	b2 = int32(b1)
	fmt.Println(b2)

	//type of a variable
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", b2)

	
}

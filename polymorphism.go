package main 

import "fmt"

type person struct{
	firstname string
}

type agent struct{
	man person
	age int
}

func(p person) speak(){
	fmt.Println(p.firstname, "says hi")
}

func(ag agent) speak(){
	fmt.Println(ag.man.firstname, "says hi")
}

type human interface{
	speak()
}

func saySomething(h human){
	h.speak()
}

func main(){
	p := person{firstname: "John"}
	a := agent{man: person{firstname: "Jane"}, age: 30}
	saySomething(p)
	saySomething(a)
}
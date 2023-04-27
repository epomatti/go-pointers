package main

import "fmt"

func main() {
	person := initPerson()
	fmt.Printf("main --> %p\n", person)
	fmt.Println(*person)
}

type person struct {
	name string
	age  int
}

func initPerson() *person {
	m := person{name: "noname", age: 50}
	fmt.Printf("initPerson --> %p\n", &m)
	return &m
}

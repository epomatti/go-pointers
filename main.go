package main

import "fmt"

func main() {
	i, j := 42, 2701

	fmt.Println("Values: ", i, j)      // Access the values
	fmt.Println("Addresses: ", &i, &j) // Access the addresses

	p := &i // this is now a pointer
	fmt.Println("Address via pointer: ", p)
	fmt.Println("Value via pointer (dereference): ", *p) // prints the value (dereferencing)

	*p = 21 // this will change the memory
	fmt.Println("Value via pointer: ", i)

	p = &j
	*p = *p / 37
	fmt.Println("Changes the value of J: ", j)

	example()
}

func example() {
	fmt.Println("--- Example ---")

	a := 4
	squareVal(a)
	squareAddress(&a)

	fmt.Println(a, &a)
}

func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}

func squareAddress(p *int) {
	*p *= *p
	fmt.Println(p, *p)
}

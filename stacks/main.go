package main

import "fmt"

func main() {
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

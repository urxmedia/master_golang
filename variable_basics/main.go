package main

import "fmt"

func main() {
	// int age = 10;	//C++ way
	var age int = 30 // data type specified explicitly
	fmt.Println("Age:", age)

	var name = "Utsav" // data type automatically inferred using the value
	fmt.Println("Your name is:", name)

	var n = "Utsav"
	_ = n // Blank identifier

	s := "Learning Golang!!!" // short-hand declaration - works only in block scope
	fmt.Println(s)

	// reassign short-hand value
	s = "ABCD"
	fmt.Println(s)

	// multiple declaration
	car, cost := "BMW", 100000
	fmt.Println(car, cost)

	// var-block declaration
	var (
		a int
		b string
		c bool
	)
	fmt.Println(a, b, c)

	// single-line same-type decalaration
	var ab, bc, cd int
	fmt.Println(ab, bc, cd)

	// multiple assignments
	var i, j int
	i, j = 5, 6
	fmt.Println(i, j)
	_, _ = i, j
	j, i = i, j // swapping variables using multi-assignment
	fmt.Println(i, j)

	// expression using short declaration
	sum := 5 + 2.4
	fmt.Println(sum)
}

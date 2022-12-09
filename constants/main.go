package main

import (
	"fmt"
)

func main() {
	const a int = 5 // compiles even when not used

	const b float64 = 3.245
	fmt.Println(b)

	const x, y = 2, 3
	fmt.Println(x, y)

	// Grouped constants with same value
	const (
		x1 = 1
		x2
		x3
	)
	fmt.Println(x1, x2, x3)

	// Constants rules
	// 1. cannot change value
	const temp int = 500
	//temp = 78

	// 2. cannot assign constants at runtime; they only work in compile-time
	//const x4 = math.Pow(2, 3)

	// 3. cannot use variables to initialise constants
	//tv := 5
	//const tc = tv

	// 4. can use builtin functions to initialise constant
	//s := "hello"
	//const sc = len(s)
	const sc = len("hello")

	// constant expressions
	const xx float32 = 4.5
	const yy = 5.6
	const zz = xx * yy

	const bb = 5 > 10
	fmt.Println(bb)

	// operations are allowed on untyped constants - not typed constants
	const ab = 89
	const bc = 90
	const abc = ab * bc
	fmt.Println(abc)

	// iota - generates untyped integer constants serially
	const (
		abcx = iota
		bcd
		cde
	)
	fmt.Println(abcx, bcd, cde)
}

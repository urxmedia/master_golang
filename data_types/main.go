package main

import (
	"fmt"
	"math"
)

func main() {
	// Minimum and Maximum vaues of numeric types
	// uint is alias for uint32 or uint64 based on platform
	// int is alias for int32 or int64 based on platform
	fmt.Printf("int8:\t[%d, %d]\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16:\t[%d, %d]\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32:\t[%d, %d]\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64:\t[%d, %d]\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("uint8:\t[%d, %d]\n", 0, math.MaxUint8)
	fmt.Printf("uint16:\t[%d, %d]\n", 0, math.MaxUint16)
	fmt.Printf("uint32:\t[%d, %d]\n", 0, math.MaxUint32)
	fmt.Printf("uint64:\t[%d, %d]\n", 0, uint64(math.MaxUint64))
	fmt.Printf("float32:[%f, %f]\n", math.SmallestNonzeroFloat32, float32(math.MaxFloat32))
	fmt.Printf("float64:[%f, %f]\n", math.SmallestNonzeroFloat64, float64(math.MaxFloat64))

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// complex numbers - complex64 and complex128
	fmt.Println("Complex numbers")
	var c1 complex64 = complex(0, 7)
	var c2 complex128 = complex(4.6, 9.67)
	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// RUNES - characters
	// rune - alias for int32
	fmt.Println("Rune")
	var r rune = '*'
	fmt.Printf("Type of r: %T\n", r)
	fmt.Printf("ASCII decimal code of r: %d\n", r)
	fmt.Printf("Hexadecimal code of r: %x\n", r)
	fmt.Printf("Octal code of r: %o\n", r)
	fmt.Printf("Binary code of r: %b\n", r)
	fmt.Printf("r: %c\n", r)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// byte - alias of uint8
	fmt.Println("byte")
	var b byte = '9'
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("ASCII decimal code of b: %d\n", b)
	fmt.Printf("Hexadecimal code of b: %x\n", b)
	fmt.Printf("Octal code of b: %o\n", b)
	fmt.Printf("Binary code of b: %b\n", b)
	fmt.Printf("b: %c\n", b)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	var bool1 bool = false
	var bool2 bool = true
	fmt.Println("bool:", bool1, bool2)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	//  string - unicode characters written in double-quotes
	var s string = "ABCD"
	fmt.Printf("string --> type: %T, value: %s\n", s, s)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// array - sequence of elements of same type with fixed-length
	var arr = [4]int{2, 0, -3, 89}
	fmt.Printf("arr --> type: %T, value: %v\n", arr, arr)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// slice - sequence of elements of same type with variable-length
	var arrSlice = []int{2, 0, -3, 89, -56, 77}
	fmt.Printf("arrSlice --> type: %T, value: %v\n", arrSlice, arrSlice)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// map - sequence of unordered key-value pairs
	m := map[string]int{
		"A": int('A'),
		"B": int('B'),
	}
	fmt.Printf("Map --> Type: %T, Value: %v\n", m, m)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// struct - sequence of named elements called fields - corresponds to class in OOP
	type Person struct {
		FirstName string
		LastName  string
	}
	var p Person
	fmt.Printf("struct --> Type: %T, Value: %+v\n", p, p)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// pointer - variable that holds memory address of another vaiable
	var x int = 6
	ptr := &x
	fmt.Printf("pointer --> Type: %T, value: %p, original value: %v\n", ptr, ptr, *ptr)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	// function
	var f = func() {}
	fmt.Printf("Function --> Type: %T\n", f)

	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	fmt.Println("more complex types: interface, channel")
}

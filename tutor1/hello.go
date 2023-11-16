package main

import (
	"fmt"
)
// list of the data types
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32

// float32 float64
// complex64 complex128
// var c complex64 = 5+5i

// main is the entry point for the application.


func main() {
	var t bool = true
	var f bool

	fmt.Println("t is ", t)
	fmt.Println("f is ", f)

	var age int = 40
	var favNum float64 = 1.6180339

	var complex complex128 = 5+5i

	// What is rune?
	//	- rune is an alias for int32
	//	- represents a Unicode code point
	//	- when you iterate over a string in Go, you get back runes
	var r rune = 10 // rune is an alias for int32

	fmt.Println("age is ", age, "favNum is ", favNum)
	fmt.Println("complex128 is ", complex)
	fmt.Println("rune is ", r)

	var myName string = "อนุสรณ์ ใจแก้ว"
	fmt.Println(myName + " is a robot")
	fmt.Println(myName[3])
	fmt.Println("length of myName is ", len(myName))

	// Arrays variables
	var arry5[5] float64 // array of 5 float64
	arry5[0] = 98.7 // assign value to array
	arry5[1] = 93.2 // assign value to array
	arry5[2] = 77.2 // assign value to array
	arry5[3] = 83.7 // assign value to array
	arry5[4] = 88.2 // assign value to array
	fmt.Println(arry5)
	fmt.Println("length of arry5 is ", len(arry5))
	fmt.Println("arry5[3] is ", arry5[3])

	// another way to declare array
	arry3 := [3]float64 { 98, 93, 77 } 
	fmt.Println(arry3)

	// slice is a segment of an array
	// slice is a reference type
	// slice is a pointer to an array
	// slice is a length and a capacity
	// slice is a dynamic size
	// slice is a flexible view into the elements of an array

	var s []float64 = arry5[0:2] // slice of an array
	fmt.Println(s)

	// struct is a collection of fields
	// struct is a user-defined type
	// struct is a composite type
	// struct is a collection of fields
	// struct is a blueprint

	type Person struct {
		name string
		age int
	}

	var p Person
	p.name = "Anusorn"
	p.age = 40
	fmt.Println(p)

	// pointer is a variable that stores the memory address of another variable
	// pointer is a reference type
	// pointer is a value that points to the address of a value

	var x int = 5
	var xPtr *int = &x
	fmt.Println(xPtr) // print the memory address of x

}
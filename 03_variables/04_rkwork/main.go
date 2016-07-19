package main

import "fmt"

// Understand variables.

// Declare using var
// double quotes give the string value
var a = "Hi Hello"

// single quote gives the ASCII value
var i = 'm'

// The following will result in an error as single quote can accept only
// one string
// var j = 'how are you'

// Tthe following is declared here, but initialized in func main
// It is in package scope and not function scope
var d string

// Shorthand:
//    Can only be used in Func.

func main() {
	fmt.Printf("This is i: %v  \n", i)
	d = "This is a new string"
	fmt.Print("hi, lets learn some variables.")
	// in func, please use :=
	b := "how are you?"
	fmt.Print(b)
	fmt.Print(a)

	// var can be used in function as well.
	var c = "new Variable \n \n \n"
	fmt.Print(c)

	var e string
	var f int
	var g float64
	var h bool

	fmt.Printf("e is : %v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
	fmt.Println()
}

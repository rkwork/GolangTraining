package main

import "fmt"

const p = "death & taxes"
const q string = "hello"

func main() {
	// scope plays a role here
	// priority goes to local scope
	//const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

// a CONSTANT is a simple unchanging value

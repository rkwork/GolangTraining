package main

import "fmt"

func main() {

	a := 43
	b := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Println("b's memory address - ", &b)
	fmt.Printf("%d \n", &a)
}

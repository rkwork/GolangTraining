package main

import (
	"fmt"

	"github.com/rkwork/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
	// Remember, var names begining with upper case letters
	// can only be visible outside.
	fmt.Println(vis.YourName)
}

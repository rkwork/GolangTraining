package main

import (
	"fmt"

	"github.com/GoesToEleven/GolangTraining/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	// uncommenting the below line will result in error as reverse() begins with a lower case letter and hence is not exported.
	//fmt.Println(stringutil.reverse("!oG ,olleH"))

	fmt.Println(stringutil.MyName)

	// uncommenting following line will result in error.
	//fmt.Println(stringutil.nyName)
}

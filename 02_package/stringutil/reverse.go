// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
// Function name with Capital Letter refers to exported function.
// Function name with lower case letter regers to non-exported function.
func Reverse(s string) string {
	return reverseTwo(s)
}

func reverse(s string) string {
	return reverseTwo(s)
}

/*
go build
	go build reverse.go reverseTwo.go
 	won't produce an output file.

go install
 	will place the package inside the pkg directory of the workspace.
*/

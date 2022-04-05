// A package roughly maps to a namespace in C#, or a module in Node.
package main

/*
	Importing a package is like including a header file in C/C++, or importing
	a module in Node. This is how one pulls in dependencies from outside of
	the current package.
*/
import (
	"fmt" //fmt is a package that provides a set of functions for formatting and printing data.
)

// The entrypoint of a package is the main function.
func main() {
	fmt.Println("Hello, World!")
}

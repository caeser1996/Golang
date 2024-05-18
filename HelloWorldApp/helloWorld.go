// Type: Go Source Code
// Description: A simple program that prints "Hello, World!" to the console
package main

// Why the package name main?
// Main is special package name which tell main is the entrypoint of the program.
// go run command will look for main package and main function.
// In production I will run go build to build the binary and run the binary.
// go run is used for running the code without building the binary.
// Once run the go build it will create the binary file with the package name.

// Import the fmt package
// fmt is short for format and is used for formatting input and output
import "fmt"

func main() {
	// Print "Hello, World!" to the console
	// Println is short for print line
	fmt.Println("Hello, World!")
}
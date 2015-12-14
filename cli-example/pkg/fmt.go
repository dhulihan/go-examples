package main

import "fmt"

func main() {
	// Strings
	fmt.Printf("\"Hello\" %%T: %T \n", "Hello")
	fmt.Printf("\"Hello\" %%q: %q \n", "Hello")
	fmt.Printf("\"Hello\" %%s: %s \n", "Hello")
	fmt.Printf("\"Hello\" %%x: %x \n", "Hello")
	fmt.Printf("\"Hello\" %%X: %X \n", "Hello")
}
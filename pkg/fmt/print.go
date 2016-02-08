package main

import "fmt"

type User struct {
	name string
	email string
}

func main() {
	// Strings
	fmt.Print("fmt.Print ", "delim", "example", "\n")
	fmt.Println("fmt.Println", "delim", "example")

	fmt.Printf("\"Hello\" %%T: %T \n", "Hello")
	fmt.Printf("\"Hello\" %%q: %q \n", "Hello")
	fmt.Printf("\"Hello\" %%s: %s \n", "Hello")
	fmt.Printf("\"Hello\" %%x: %x \n", "Hello")
	fmt.Printf("\"Hello\" %%X: %X \n", "Hello")

	fmt.Printf("Generic struct: %v \n", User{"Bob Jones", "bob@example.com"})
}
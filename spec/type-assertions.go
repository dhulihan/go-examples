package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = 7
	i, ok := x.(int) // check if i is an int

	fmt.Printf("Subject is type %s\n", reflect.TypeOf(x))

	fmt.Println("Checking if int...")
	fmt.Printf("i: %#v, ok: %#v\n\n", i, ok)

	fmt.Println("Checking if string...")
	j, ok := x.(string) // check if i is a string
	fmt.Printf("j: %#v, ok: %#v\n", j, ok)
}
package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", os.Getenv("USER"))
	
	for _, envvar := range os.Environ() {
		fmt.Printf("%s\n", envvar)
	}
}
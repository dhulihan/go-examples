package main

import (
	"fmt"
	"os"
)

func main() {
	// create closure and call
	func(){ fmt.Println("Hello!") }()
	func(name string){ fmt.Println("Nice to meet you", name) }(os.Getenv("USER")) // 
}

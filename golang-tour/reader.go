package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, my wonderful friend!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		if err == io.EOF {
			break
		}
	}
}

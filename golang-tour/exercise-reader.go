package main

// Assignment: Implement a Reader type that emits an infinite stream of the ASCII character 'A'.


import (
	"fmt"
	// "golang.org/x/tour/reader"
	"io"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m *MyReader) Read(a []byte) (int, error) {
		char := 'A'
		counter := 0
		for pos := range a {
			a[pos] = byte(char)
			counter++
		}
		err := io.EOF
		return counter, err
}

func main() {
	chars := []byte {'A', 'B', 'C', 'D', 'E', 'F'}
	fmt.Printf("chars before: %#v\n", chars)

	reader := MyReader{}
	n, err := reader.Read(chars)
	fmt.Printf("chars: %q\n", chars)
	fmt.Printf("n: %v \n", n)
	fmt.Printf("err: %v \n", err.Error())
	// reader.Validate(MyReader{})
}

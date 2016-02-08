package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	var s [][]uint8
	for y := 0; y < dy; y++ {
		var xs []uint8
		for x := 0; x < dx; x++ {
			xs = append(xs, uint8(x))
		}
		s = append(s, xs)
	}
	return s
}

func main() {
	fmt.Println(Pic(5, 5))
}

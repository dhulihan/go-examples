package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("You can't get the square root a negative number (%v)", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	vals_to_test := []float64{2, -2}
	for _, val := range vals_to_test {
		sqrt, err := Sqrt(val)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(sqrt)
		}
	}
}

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) String() string {
    return fmt.Sprintln(v.X, v.Y)
}

func main(){
	v := &Vertex{92, -13}
	fmt.Println(v.String());
}
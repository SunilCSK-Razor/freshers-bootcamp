package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func AbsFunc(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex4{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex4{4, 3}
	fmt.Println(p)
	fmt.Println(p.Abs())
	fmt.Println((*p).Abs())
	fmt.Println(AbsFunc(*p))
}

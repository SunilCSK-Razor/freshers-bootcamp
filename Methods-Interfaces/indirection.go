package main

import "fmt"

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex3{3, 4}
	v.Scale(2)
	fmt.Println(v)
	ScaleFunc(&v, 10)
	p := &Vertex3{4, 3}
	fmt.Println(p)
	p.Scale(3)
	fmt.Println(p)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

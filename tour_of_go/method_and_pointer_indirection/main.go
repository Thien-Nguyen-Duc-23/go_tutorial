package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	fmt.Println("-----------------")
	v.Scale(2)
	fmt.Println(v)
	fmt.Println("-----------------")
	ScaleFunc(&v, 10)
	fmt.Println(v)
	fmt.Println("-----------------")

	p := &Vertex{4, 3}
	fmt.Println(p)
	fmt.Println("-----------------")

	p.Scale(3)
	fmt.Println(p)
	fmt.Println("-----------------")

	ScaleFunc(p, 8)
	fmt.Println(p)
	fmt.Println("-----------------")

	fmt.Println(v, p)

}

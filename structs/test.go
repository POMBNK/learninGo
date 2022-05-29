package main

import "fmt"

const pi float64 = 3.1415

type Circle struct {
	x float64
	y float64
	r float64
}

func Area(obj Circle) float64 {
	return pi * obj.r * obj.r
}

func (p Circle) Area() float64 {
	return pi * p.r * p.r
}

func main() {
	krug := Circle{
		x: 2,
		y: 2,
		r: 1,
	}
	krug.r = 5
	fmt.Println(krug)
	fmt.Println(Area(krug))
	fmt.Println(krug.Area())

}

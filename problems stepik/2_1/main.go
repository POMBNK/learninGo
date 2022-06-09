package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a float64
		b float64
	)
	fmt.Scanln(&a, &b)
	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)
}

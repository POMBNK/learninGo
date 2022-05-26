package main

import "fmt"

func main() {
	res := minimumFromFour()
	fmt.Println(res)
}

func minimumFromFour() int {
	var (
		min int
		a   int
		b   int
		c   int
		d   int
	)
	fmt.Scan(&a, &b, &c, &d)
	if a < b && a < c && a < d {
		min = a
	} else if b < a && b < c && b < d {
		min = b
	} else if c < a && c < b && c < d {
		min = c
	} else {
		min = d
	}
	return min
}

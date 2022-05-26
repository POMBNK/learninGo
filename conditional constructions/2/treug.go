package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
	)
	fmt.Scan(&a, &b, &c)
	if a*a+b*b == c^2 {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

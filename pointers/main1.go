package main

import "fmt"

func test(x1 *int, x2 *int) {
	res := *x1 * *x2
	fmt.Println(res)
}

func main() {
	x1 := 2
	x2 := 4
	test(&x1, &x2)
}

package main

import "fmt"

func dec(wrapper func(a, b int) int, s string) int {
	result := wrapper(10, 2)
	fmt.Println(s)
	return result
}

func main() {
	add := func(a, b int) int {
		return a + b
	}

	someRes := dec(add, "add")
	fmt.Println(someRes)
}

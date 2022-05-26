package main

import "fmt"

// Попытка в декоратор или колбэк
func dec(wrapper func(a, b int) int, s string) int {
	result := wrapper(3, 2)
	fmt.Println(s)
	return result

}

func main() {
	multiply := func(a, b int) int {
		return a * b
	}
	someResult := dec(multiply, "multiply")
	fmt.Println(someResult)
}

package main

import "fmt"

func main() {
	var (
		n   int
		num int
	)
	fmt.Scan(&n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		slice[i] = num
	}
	for i := 0; i < len(slice); i++ {
		if i%2 == 0 {
			fmt.Print(slice[i], " ")
		}
	}
}

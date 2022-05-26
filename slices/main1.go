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
	fmt.Println(slice[3])
}

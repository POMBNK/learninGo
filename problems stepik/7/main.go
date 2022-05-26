package main

import "fmt"

// find count of zeros in nums from keyboard
func main() {
	var (
		n     int
		num   int
		count int
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num == 0 {
			count++
		}
	}
	fmt.Println(count)
}

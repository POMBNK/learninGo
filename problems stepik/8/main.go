package main

import "fmt"

func main() {
	var (
		n     int
		num   int
		min   int
		count int
	)
	fmt.Scan(&n)
	fmt.Scan(&num)
	min = num
	count = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&num)
		if num == min {
			count++
		}
		if min > num {
			min = num
			count = 1
		}
	}
	fmt.Println(count)
}

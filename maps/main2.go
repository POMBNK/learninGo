package main

import (
	"fmt"
	"time"
)

//work function is already given
func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}

func main() {
	var num int
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		if _, ok := m[num]; ok {
			fmt.Print(m[num], " ")
		} else {
			m[num] = work(num)
			fmt.Print(m[num], " ")
		}
	}

}

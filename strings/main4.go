package main

import "fmt"

func NotEven(word string) string {
	res := ""
	fmt.Scanln(&word)
	rs := []rune(word)
	for i := 1; i <= len(rs)-1; i++ {
		if i%2 == 1 {
			res += string(rs[i])
		}
	}
	return res
}

func main() {
	fmt.Println(NotEven("ihgewlqlkot"))
}

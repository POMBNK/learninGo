package main

import "fmt"

func main() {
	var word string
	fmt.Scanln(&word)
	rs := []rune(word)
	rev := ""
	length := len(rs)
	for i := length - 1; i >= 0; i-- {
		rev += string(rs[i])
	}
	if word == rev {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

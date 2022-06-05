package main

import (
	"fmt"
	"strings"
)

func delReap(word string) string {
	rs := []rune(word)
	res := ""
	for i := 0; i < len(word); i++ {
		if strings.Count(word, string(rs[i])) > 1 {
			continue
		} else {
			res += string(rs[i])
		}
	}
	return res
}

func main() {
	fmt.Println(delReap("zaabcbd")) // zcd
}

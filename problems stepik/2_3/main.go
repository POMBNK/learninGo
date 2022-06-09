package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		s   string
		max int
	)
	fmt.Scanln(&s)
	strings.Split(s, "")
	max, _ = strconv.Atoi(string(s[0]))
	for i := 0; i < len(s)-1; i++ {
		conv, _ := strconv.Atoi(string(s[i]))
		if conv > max {
			max = conv
		}
	}
	fmt.Println(max)
}

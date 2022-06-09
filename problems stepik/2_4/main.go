package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		s string
	)
	fmt.Scanln(&s)
	strings.Split(s, "")
	res := ""
	for i := 0; i < len(s); i++ {
		conv, _ := strconv.Atoi(string(s[i]))
		kv := conv * conv
		res += strconv.Itoa(kv)
	}
	fmt.Println(res)
}

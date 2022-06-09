package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(strings.Join(strings.Split(s, ""), "*"))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		X string
		S string
	)
	fmt.Scanln(&X)
	fmt.Scanln(&S)
	fmt.Println(strings.Index(X, S))
}

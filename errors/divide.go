package main

import (
	"fmt"
)

// func divide(a,b int) (int,error) already exist

func main() {
	var (
		a int
		b int
	)
	fmt.Scanln(&a, &b)
	c, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}

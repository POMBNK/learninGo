package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	res := ""
	for i := 0; i <= len(s)-2; i++ {
		res += string(s[i]) + "*"
	}
	res += string(s[len(s)-1])
	fmt.Println(res)
}

//the worst hardcode, some better at main1.go

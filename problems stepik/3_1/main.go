package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func adding(s1, s2 string) int64 {
	rs1 := []rune(s1)
	rs2 := []rune(s2)
	num1 := ""
	num2 := ""
	for i := 0; i < len(rs1); i++ {
		if unicode.IsDigit(rs1[i]) {
			num1 += string(rs1[i])
		}
		continue
	}
	res1, _ := strconv.Atoi(num1)
	for i := 0; i < len(rs2); i++ {
		if unicode.IsDigit(rs2[i]) {
			num2 += string(rs2[i])
		}
		continue
	}
	res2, _ := strconv.Atoi(num2)
	return int64(res1 + res2)
}

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}

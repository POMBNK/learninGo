package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)
	if string(unicode.ToUpper(rs[0])) == string(rs[0]) && rs[len(rs)-3] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

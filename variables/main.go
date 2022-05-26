package main

import "fmt"

const (
	OK = iota
	CANCELLED
	UNKNOWN
)

func ErrorMessageToCode(msg string) int {
	switch msg{
		case "OK":
			return OK
		case "CANCELLED":
			return CANCELLED
		case "UNKNOWN":
			return UNKNOWN
		default: 
			return UNKNOWN 
	}
}


func main() {
	fmt.Println(ErrorMessageToCode("OK"))
}

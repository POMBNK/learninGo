package main

const (
	OkMsg        = "OK"
	CancelledMsg = "CANCELLED"
)

const (
	OkCode = iota
	CancelledCode
	UnknownCode
)

func ErrorMessageToCode(msg string) int {
	switch msg {
	case OkMsg:
		return OkCode
	case CancelledMsg:
		return CancelledCode
	}

	return UnknownCode
}

func main() {
	ErrorMessageToCode("OK")
}

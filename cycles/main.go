package main

/* Программа должна вывести цифры,
которые имеются в обоих числах, через пробел.
Цифры выводятся в порядке их нахождения в первом числе. */

import "fmt"

func main() {
	var (
		a string
		b string
	)
	fmt.Scan(&a, &b)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if b[j] == a[i] {
				fmt.Print(string(b[j]), " ")
			}

		}

	}
}

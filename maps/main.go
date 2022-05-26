package main

import "fmt"

func UniqueUserIDs(userIDs []int64) []int64 {
	m := make(map[int64]struct{})
	ids := make([]int64, 0)

	for _, elem := range userIDs {
		fmt.Println(m)
		if _, ok := m[elem]; ok {
			continue
		}

		ids = append(ids, elem)
		fmt.Println(ids)
		m[elem] = struct{}{}
	}
	return ids
}

func main() {
	UniqueUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88})
}

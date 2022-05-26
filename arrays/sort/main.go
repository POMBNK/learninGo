package main

import (
	"fmt"
	"sort"
)

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	sort.Slice(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})
	for i := 1; i < len(userIDs); i++ {
		if userIDs[i] == userIDs[i-1] {
			userIDs = append(userIDs[:i], userIDs[i+1:]...)
			i -= 1
		}
	}
	return userIDs
}

func main() {
	fmt.Println(UniqueSortedUserIDs([]int64{55, 55}))
	fmt.Println(UniqueSortedUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88, 2, 2}))
}

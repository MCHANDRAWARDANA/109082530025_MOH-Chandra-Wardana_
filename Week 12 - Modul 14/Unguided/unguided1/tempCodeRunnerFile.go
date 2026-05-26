package main

import (
	"fmt"
)

func selectionSortAscending(arr []int, n int) {
	for i := 1; i <= n-1; i++ {
		idxMin := i - 1
		for j := i; j < n; j++ {
			if arr[idxMin] > arr[j] {
				idxMin = j
			}
		}
		t := arr[idxMin]
		arr[idxMin] = arr[i-1]
		arr[i-1] = t
	}
}

func main() {
	var n, m int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for k := 0; k < n; k++ {
		fmt.Scan(&m)
		rumah := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&rumah[i])
		}

		selectionSortAscending(rumah, m)

		for i := 0; i < m; i++ {
			fmt.Printf("%d", rumah[i])
			if i < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
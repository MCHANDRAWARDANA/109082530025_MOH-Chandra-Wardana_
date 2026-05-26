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

func selectionSortDescending(arr []int, n int) {
	for i := 1; i <= n-1; i++ {
		idxMax := i - 1
		for j := i; j < n; j++ {
			if arr[idxMax] < arr[j] {
				idxMax = j
			}
		}
		t := arr[idxMax]
		arr[idxMax] = arr[i-1]
		arr[i-1] = t
	}
}

func main() {
	var n, m, nomor int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for k := 0; k < n; k++ {
		fmt.Scan(&m)
		var ganjil []int
		var genap []int

		for i := 0; i < m; i++ {
			fmt.Scan(&nomor)
			if nomor%2 != 0 {
				ganjil = append(ganjil, nomor)
			} else {
				genap = append(genap, nomor)
			}
		}

		selectionSortAscending(ganjil, len(ganjil))
		selectionSortDescending(genap, len(genap))

		var hasil []string
		for _, v := range ganjil {
			hasil = append(hasil, fmt.Sprintf("%d", v))
		}
		for _, v := range genap {
			hasil = append(hasil, fmt.Sprintf("%d", v))
		}

		for i, v := range hasil {
			fmt.Print(v)
			if i < len(hasil)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
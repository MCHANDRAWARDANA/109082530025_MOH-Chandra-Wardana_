package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

// Selection Sort
func selectionSort(T *arrInt, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if T[j] < T[idxMin] {
				idxMin = j
			}
		}

		temp = T[i]
		T[i] = T[idxMin]
		T[idxMin] = temp
	}
}

// Function median
func median(T arrInt, n int) float64 {
	if n%2 == 1 {
		return float64(T[n/2])
	}

	return float64(T[(n/2)-1]+T[n/2]) / 2
}

func main() {
	var T arrInt
	var x int
	var n int = 0

	fmt.Println("Input data masukan :")

	fmt.Scan(&x)

	for x != -5313541 && n < NMAX {

		if x == 0 {

			// urutkan data
			selectionSort(&T, n)

			// tampilkan median
			fmt.Println("Median :")
			fmt.Println(median(T, n))

		} else {

			// simpan data ke array
			T[n] = x
			n++
		}

		fmt.Scan(&x)
	}
}
package main

import "fmt"

func main() {
	var x int
	var totalSuara int
	var suaraSah int

	var calon [21]int

	for {
		fmt.Scan(&x)

		totalSuara++

		if x == 0 {
			break
		}

		if x >= 1 && x <= 20 {
			calon[x]++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk:", totalSuara)
	fmt.Println("Suara sah:", suaraSah)

	for i := 1; i <= 20; i++ {
		if calon[i] > 0 {
			fmt.Printf("%d: %d\n", i, calon[i])
		}
	}
}
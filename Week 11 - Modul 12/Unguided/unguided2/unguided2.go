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

	ketua := 1
	wakil := 1

	for i := 2; i <= 20; i++ {
		if calon[i] > calon[ketua] {
			ketua = i
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			wakil = i
			break
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if calon[i] > calon[wakil] {
				wakil = i
			} else if calon[i] == calon[wakil] && i < wakil {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", totalSuara)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
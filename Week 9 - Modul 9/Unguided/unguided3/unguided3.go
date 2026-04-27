package main

import "fmt"

func main() {
	klubA := "MU"
	klubB := "Inter"

	skorA := []int{2, 1, 2, 0, 3, 5, 2}
	skorB := []int{1, 2, 2, 1, 1, 2, 3}

	var pemenang []string

	for i := 0; i < len(skorA); i++ {
		if skorA[i] > skorB[i] {
			pemenang = append(pemenang, klubA)
		} else if skorB[i] > skorA[i] {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
	}

	fmt.Println("Hasil pertandingan:")
	for i, hasil := range pemenang {
		fmt.Println("Pertandingan", i+1, ":", hasil)
	}
}
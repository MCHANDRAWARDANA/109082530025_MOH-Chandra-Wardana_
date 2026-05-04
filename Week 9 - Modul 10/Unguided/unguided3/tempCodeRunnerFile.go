package main

import "fmt"

func main() {
	var n int
	var berat [100]float64
	var min, max, total float64

	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min = berat[0]
	max = berat[0]

	for i := 0; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
		total += berat[i]
	}

	rata := total / float64(n)

	fmt.Println("Berat minimum:", min)
	fmt.Println("Berat maksimum:", max)
	fmt.Println("Rata-rata:", rata)
}
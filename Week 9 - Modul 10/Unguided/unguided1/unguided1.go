package main

import "fmt"

func main() {
	var n int
	var berat [1000]float64
	var min, max float64

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min = berat[0]
	max = berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}
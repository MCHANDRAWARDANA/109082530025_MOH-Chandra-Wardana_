package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{10, 20, 30, 40, 50}

	fmt.Println("Isi array:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Print("\nIndeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}

	x := 2
	fmt.Print("\nKelipatan indeks ", x, ": ")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	hapus := 2
	arr = append(arr[:hapus], arr[hapus+1:]...)
	fmt.Println("\nArray setelah hapus indeks 2:", arr)

	total := 0
	for _, v := range arr {
		total += v
	}
	rata := float64(total) / float64(len(arr))
	fmt.Println("Rata-rata:", rata)

	var jumlah float64
	for _, v := range arr {
		jumlah += math.Pow(float64(v)-rata, 2)
	}
	std := math.Sqrt(jumlah / float64(len(arr)))
	fmt.Println("Standar deviasi:", std)
}
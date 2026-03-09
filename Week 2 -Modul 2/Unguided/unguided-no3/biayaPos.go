package main

import "fmt"

func main() {
	var berat, kg, sisa, biaya int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scanln(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biaya = kg * 10000

	if kg > 10 {
		sisa = 0
	} else {
		if sisa > 500 {
			biaya = biaya + (sisa * 5)
		} else {
			biaya = biaya + (sisa * 15)
		}
	}

	fmt.Println("Detail berat:", kg, "kg dan", sisa, "gram")
	fmt.Println("Biaya per kg : Rp10000")
	fmt.Println("Total biaya pengiriman: Rp", biaya)
}
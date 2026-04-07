package main

import "fmt"

// Menentukan maksimal hari
func tentukanHari(lamaHari int, tujuan string) int {
	if tujuan == "domestik" && lamaHari > 3 {
		return 3
	}
	if tujuan == "mancanegara" && lamaHari > 8 {
		return 8
	}
	return lamaHari
}

// Menghitung total biaya
func hitungBiaya(jumlah int, lamaHari int, tujuan string) int {
	hari := tentukanHari(lamaHari, tujuan)

	biayaMakan := 3 * 50000
	biayaPenginapan := 250000
	uangSaku := 300000

	biayaHarian := biayaMakan + biayaPenginapan + uangSaku

	if tujuan == "mancanegara" {
		biayaHarian = int(float64(biayaHarian) * 1.5)
	}

	return jumlah * hari * biayaHarian
}

func main() {
	var jumlah, lamaHari int
	var tujuan string

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan tujuan study tour : ")
	fmt.Scan(&tujuan)

	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lamaHari)

	total := hitungBiaya(jumlah, lamaHari, tujuan)

	fmt.Println(total)
}

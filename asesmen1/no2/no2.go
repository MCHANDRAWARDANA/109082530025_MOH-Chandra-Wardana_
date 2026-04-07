package main

import "fmt"

func tentukanHari(lamaHari int, tujuan string) int {
	if tujuan == "domestik" {
		if lamaHari > 3 {
			return 3
		}
	} else if tujuan == "mancanegara" {
		if lamaHari > 8 {
			return 8
		}
	}
	return lamaHari
}

func biayaPerHariDomestik() float64 {
	biayaMakan := 3 * 50000
	biayaPenginapan := 250000
	uangSaku := 300000

	return float64(biayaMakan + biayaPenginapan + uangSaku)
}

func hitungTotalBiaya(jumlah int, lamaHari int, tujuan string, totalBiaya *float64) {
	hari := tentukanHari(lamaHari, tujuan)
	biayaHarian := biayaPerHariDomestik()

	if tujuan == "mancanegara" {
		biayaHarian = biayaHarian * 1.5
	}

	*totalBiaya = float64(jumlah) * float64(hari) * biayaHarian
}

func main() {
	var jumlah int
	var lamaHari int
	var tujuan string
	var totalBiaya float64

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lamaHari)

	hitungTotalBiaya(jumlah, lamaHari, tujuan, &totalBiaya)

	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp %.0f\n", totalBiaya)
}

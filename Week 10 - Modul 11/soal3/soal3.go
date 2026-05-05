package main

import (
	"fmt"
	"math"
)

const nProv int = 10

type NamaProv [10]string
type PopProv [10]int
type TumbuhProv [10]float64

func inputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func provinsiTercepat(tumbuh TumbuhProv) int {
	idxMax := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxMax] {
			idxMax = i
		}
	}
	return idxMax
}

func prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksiPop := math.Round((tumbuh[i]+1) * float64(pop[i]))
			fmt.Printf("%s %d\n", prov[i], int(prediksiPop))
		}
	}
}

func indeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	inputData(&prov, &pop, &tumbuh)

	var namaCari string
	fmt.Scan(&namaCari)

	idx := provinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idx])

	idxCari := indeksProvinsi(prov, namaCari)
	fmt.Printf("\nData provinsi yang dicari : %s\n", prov[idxCari])

	prediksi(prov, pop, tumbuh)
}
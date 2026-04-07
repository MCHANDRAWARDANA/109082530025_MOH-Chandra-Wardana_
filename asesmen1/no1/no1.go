package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.14


func volume(r, t float64) float64 {
	return pi * r * r * t
}


func massa(r, t, rho float64) float64 {
	return volume(r, t) * rho
}


func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		selisih := math.Abs(m1 - m2)
		fmt.Printf("Selisih massa zat cair kiri dan massa zat cair kanan : %.0f\n", selisih)
	}
}

func main() {
	var r float64
	var tKiri, tKanan float64
	var rhoKiri, rhoKanan float64
	var mKiri, mKanan float64

	
	fmt.Print("Masukkan jari-jari alas tabung: ")
	fmt.Scan(&r)

	
	fmt.Print("\nMasukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&rhoKiri)

	
	fmt.Print("\nMasukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&rhoKanan)

	mKiri = massa(r, tKiri, rhoKiri)
	mKanan = massa(r, tKanan, rhoKanan)

	display(mKiri, mKanan)
}

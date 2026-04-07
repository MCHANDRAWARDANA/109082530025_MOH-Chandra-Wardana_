package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.14


func volume(r, t float64) float64 {
	return pi * math.Pow(r, 2) * t
}


func massa(r, t, rho float64) float64 {
	return volume(r, t) * rho
}


func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		selisih := math.Abs(m1 - m2)
		fmt.Printf("Selisih massa: %.2f\n", selisih)
	}
}

func main() {
	var r float64 

	var tKiri, tKanan float64 
	var rhoKiri, rhoKanan float64 

	var mKiri, mKanan float64 

	
	fmt.Print("Masukkan jari-jari alas tabung: ")
	fmt.Scan(&r)

	
	fmt.Print("Masukkan tinggi cairan tabung kiri: ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis cairan kiri: ")
	fmt.Scan(&rhoKiri)

	
	fmt.Print("Masukkan tinggi cairan tabung kanan: ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis cairan kanan: ")
	fmt.Scan(&rhoKanan)

	
	mKiri = massa(r, tKiri, rhoKiri)
	mKanan = massa(r, tKanan, rhoKanan)

	
	display(mKiri, mKanan)
}

package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return c * 0.8
}

func CelciusToFahrenheit(c suhu) suhu {
	return (c * 1.8) + 32
}

// Fungsi konversi ke Kelvin [cite: 6]
func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var c suhu
	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&c)

	fmt.Printf("%.1f celcius = %.1f reamur\n", c, CelciusToReamur(c))
	fmt.Printf("%.1f celcius = %.1f fahrenheit\n", c, CelciusToFahrenheit(c))
	fmt.Printf("%.2f celcius = %.2f kelvin\n", c, CelciusToKelvin(c))
}
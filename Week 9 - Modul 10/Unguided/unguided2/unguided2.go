package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Print("Masukkan jumlah ikan: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan kapasitas per wadah: ")
	fmt.Scan(&y)

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := (x + y - 1) / y
	var total float64

	fmt.Println("Total berat per wadah:")
	for i := 0; i < jumlahWadah; i++ {
		var sum float64
		for j := i * y; j < (i+1)*y && j < x; j++ {
			sum += ikan[j]
		}
		fmt.Println(sum)
		total += sum
	}

	rata := total / float64(jumlahWadah)
	fmt.Println("Rata-rata berat per wadah:", rata)
}
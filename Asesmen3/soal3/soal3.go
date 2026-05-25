package main

import "fmt"

const NMAX = 100000

// struct partai
type partai struct {
	nama  int
	suara int
}

// tipe tabPartai: array of partai dengan kapasitas NMAX
type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, x int) int {
	// mengembalikan indeks partai yang memiliki nama yang dicari
	// pada array t yang berisi n partai atau -1 apabila tidak ditemukan

	for i := 0; i < n; i++ {
		if t[i].nama == x {
			return i
		}
	}
	return -1
}

func main() {
	var t tabPartai
	var n int
	var x int
	var idx int

	// baca input sampai -1
	fmt.Println("Masukkan proses input suara :")

	for {
		fmt.Scan(&x)

		if x == -1 {
			break
		}

		idx = posisi(t, n, x)

		if idx == -1 {
			// partai baru
			t[n].nama = x
			t[n].suara = 1
			n++
		} else {
			// tambah suara
			t[idx].suara++
		}
	}

	// insertion sort descending berdasarkan jumlah suara
	for i := 1; i < n; i++ {
		temp := t[i]
		j := i - 1

		for j >= 0 && t[j].suara < temp.suara {
			t[j+1] = t[j]
			j--
		}

		t[j+1] = temp
	}

	// tampilkan hasil
	fmt.Println("\nHasil Perhitungan Suara :")

	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", t[i].nama, t[i].suara)
	}

	fmt.Println()
}
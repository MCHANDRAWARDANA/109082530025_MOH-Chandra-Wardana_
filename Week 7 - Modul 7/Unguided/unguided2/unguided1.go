package main

import "fmt"

// Mendefinisikan tipe data alias 
type angka int
type kata string

// Mendefinisikan struct Buku [cite: 17]
type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var b Buku

	// Input data buku
	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku: ")
	fmt.Scan(&b.judul)
	fmt.Print("Masukkan nama penulis: ")
	fmt.Scan(&b.penulis)
	fmt.Print("Masukkan penerbit: ")
	fmt.Scan(&b.penerbit)
	fmt.Print("Masukkan tahun terbit: ")
	fmt.Scan(&b.tahunTerbit)
	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&b.jumlahHalaman)

	// Output data buku
	fmt.Println("\nBIODATA BUKU")
	fmt.Println("Judul Buku:", b.judul)
	fmt.Println("Penulis:", b.penulis)
	fmt.Println("Penerbit:", b.penerbit)
	fmt.Println("Tahun Terbit:", b.tahunTerbit)
	fmt.Println("Jumlah Halaman :", b.jumlahHalaman)
}
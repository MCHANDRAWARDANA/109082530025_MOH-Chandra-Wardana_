package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(T arrayMahasiswa, n int, targetNIM string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == targetNIM {
			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, targetNIM string) int {
	max := -1
	for i := 0; i < n; i++ {
		if T[i].NIM == targetNIM {
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}

func main() {
	var T arrayMahasiswa
	var n int
	var targetNIM string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("Masukkan NIM yang dicari : ")
	fmt.Scan(&targetNIM)

	pertama := cariNilaiPertama(T, n, targetNIM)
	terbesar := cariNilaiTerbesar(T, n, targetNIM)

	if pertama != -1 {
		fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah: %d\n", targetNIM, pertama)
		fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah: %d\n", targetNIM, terbesar)
	} else {
		fmt.Println("Data mahasiswa tidak ditemukan.")
	}
}
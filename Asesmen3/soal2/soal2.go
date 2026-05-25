package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX = 100

type Pemain struct {
	nama   string
	gol    int
	assist int
}

type arrPemain [NMAX]Pemain

// Selection Sort descending
func selectionSort(A *arrPemain, n int) {
	var i, j, idxMax int
	var temp Pemain

	for i = 0; i < n-1; i++ {
		idxMax = i

		for j = i + 1; j < n; j++ {
			if A[j].gol > A[idxMax].gol {
				idxMax = j
			} else if A[j].gol == A[idxMax].gol {
				if A[j].assist > A[idxMax].assist {
					idxMax = j
				}
			}
		}

		temp = A[i]
		A[i] = A[idxMax]
		A[idxMax] = temp
	}
}

func main() {
	var pemain arrPemain
	var n int

	fmt.Print("Masukkan Jumlah Pemain: ")
	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan Data Input (Format: Nama Gol Assist):")
	for i := 0; i < n; i++ {
		fmt.Printf("Data ke-%d: ", i+1)
		
		if scanner.Scan() {
			line := scanner.Text()
			
			_, err := fmt.Sscanf(line, "%s %d %d", &pemain[i].nama, &pemain[i].gol, &pemain[i].assist)
			if err != nil {
				fmt.Println("Format salah! Gunakan: [Nama] [Gol] [Assist]")
				i-- 
				continue
			}
		}
	}

	// sorting
	selectionSort(&pemain, n)

	// output
	fmt.Println()
	fmt.Println("Hasil Sorting :")

	for i := 0; i < n; i++ {
		fmt.Printf("%s - Gol: %d, Assist: %d\n", pemain[i].nama, pemain[i].gol, pemain[i].assist)
	}
}
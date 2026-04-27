package main

import "fmt"

func balikanArray(arr []rune) []rune {
	n := len(arr)
	hasil := make([]rune, n)

	for i := 0; i < n; i++ {
		hasil[i] = arr[n-1-i]
	}
	return hasil
}

func palindrom(arr []rune) bool {
	balik := balikanArray(arr)

	for i := range arr {
		if arr[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	kata := []rune{'K', 'A', 'T', 'A', 'K'}

	fmt.Print("Array asli: ")
	for _, v := range kata {
		fmt.Print(string(v), " ")
	}

	balik := balikanArray(kata)

	fmt.Print("\nArray dibalik: ")
	for _, v := range balik {
		fmt.Print(string(v), " ")
	}

	if palindrom(kata) {
		fmt.Println("\nPalindrom")
	} else {
		fmt.Println("\nBukan Palindrom")
	}
}
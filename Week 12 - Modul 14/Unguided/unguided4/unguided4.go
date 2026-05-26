package main

import (
	"fmt"
)

func insertionSortAscending(arr []int, n int) {
	for i := 1; i <= n-1; i++ {
		j := i
		temp := arr[j]
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = temp
	}
}

func main() {
	var input int
	var data []int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	n := len(data)
	if n == 0 {
		return
	}

	insertionSortAscending(data, n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d", data[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	if n == 1 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarakAwal := data[1] - data[0]
	tetap := true

	for i := 1; i < n-1; i++ {
		if (data[i+1] - data[i]) != jarakAwal {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarakAwal)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
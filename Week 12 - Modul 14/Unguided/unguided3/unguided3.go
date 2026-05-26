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
	var data []int
	var input int

	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}

		if input == 0 {
			n := len(data)
			if n > 0 {
				insertionSortAscending(data, n)
				var median int
				if n%2 != 0 {
					median = data[n/2]
				} else {
					median = (data[(n/2)-1] + data[n/2]) / 2
				}
				fmt.Println(median)
			}
		} else {
			data = append(data, input)
		}
	}
}
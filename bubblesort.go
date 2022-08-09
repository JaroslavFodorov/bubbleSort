package main

import "fmt"

func main() {
	mas := []int{3, 9, 5, 2, -6, 0}
	bubbleSort(mas)
	fmt.Println(mas)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

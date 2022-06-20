package main

import "fmt"

func insertionSortAscendingOrder(array []int) {
	for j := 1; j < len(array); j++ {
		key := array[j]
		i := j - 1
		for i >= 0 && array[i] > key {
			array[i+1] = array[i]
			i--
		}
		array[i+1] = key
	}
}

func insertionSortDescendingOrder(array []int) {
	for j := 1; j < len(array); j++ {
		key := array[j]
		i := j - 1
		for i >= 0 && array[i] < key {
			array[i+1] = array[i]
			i--
		}
		array[i+1] = key
	}
}

func main() {
	array := []int{34, 2, 17, 11, 39, 4, 11, 20}
	fmt.Println("Given array:", array)
	insertionSortAscendingOrder(array)
	fmt.Println("Sorted in ascending order:", array)
	insertionSortDescendingOrder(array)
	fmt.Println("Sorted in descending order:", array)
}

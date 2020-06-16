// Rotates an array n number of indices to the left
package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	array = rotate(array, 1)
	fmt.Println(array)
}

func rotate(array []int, indices int) []int {
	indices = indices % len(array)
	left := make([]int, len(array[:indices]))
	copy(left, array[:indices])
	right := array[indices:]
	for i := 0; i < len(right); i++ {
		array[i] = right[i]
	}
	for i, j := len(right), 0; j < len(left); i, j = i+1, j+1 {
		array[i] = left[j]
	}

	return array
}

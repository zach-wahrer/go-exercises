// Reverse reverses an int array in place
package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	reverse(&array)
	fmt.Println(array)
}

func reverse(array *[5]int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

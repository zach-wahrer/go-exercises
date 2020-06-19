// Max is a variadic function that returns the maximum int from its arguments
package main

import "fmt"

func main() {
	fmt.Println(max([]int{1, 2, 3, 4}...))
	fmt.Println(max(1))
	fmt.Println(max(1, 8, 99, 2, 8, 15))
	fmt.Println(max())
}

func max(values ...int) int {
	if len(values) == 0 {
		values = []int{0}
	}
	maximum := values[0]
	for _, value := range values {
		if value > maximum {
			maximum = value
		}
	}
	return maximum
}

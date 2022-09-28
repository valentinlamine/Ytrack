package main

import "fmt"

func main() {
	fmt.Println(makerange(5, 10))
	fmt.Println(makerange(10, 5))
}

func makerange(min, max int) []int {
	tab_size := max - min
	if max < min {
		tab_size = 0
	}
	result := make([]int, tab_size)
	for i := min; i < max; i++ {
		result[i-min] = i
	}
	return result
}

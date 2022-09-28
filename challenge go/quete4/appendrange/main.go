package main

import "fmt"

func main() {
	fmt.Println(appendrange(5, 10))
	fmt.Println(appendrange(10, 5))
}

func appendrange(min, max int) []int {
	var result []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}

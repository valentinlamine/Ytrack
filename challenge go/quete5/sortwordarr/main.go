package main

import (
	"fmt"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(a []string) {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			temp := a[i]
			a[i] = a[i+1]
			a[i+1] = temp
			SortWordArr(a)
		}
	}
}

package main

import "fmt"

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(basicjoin(elems))
}

func basicjoin(elems []string) string {
	result := ""
	for _, strings := range elems {
		result += strings
	}
	return result
}

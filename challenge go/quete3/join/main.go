package main

import "fmt"

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(join(elems, ":"))
}

func join(elems []string, sep string) string {
	result := ""
	for index, strings := range elems {
		if index < len(elems)-1 {
			result += strings + sep
		} else {
			result += strings
		}
	}
	return result
}

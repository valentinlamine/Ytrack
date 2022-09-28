package main

import "fmt"

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(isnumeric, a1)
	result2 := Any(isnumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func isnumeric(str string) bool {
	for index, letter := range str {
		_ = index //pour utiliser la valeur index
		if letter < '0' || letter > '9' {
			return false
		}
	}
	return true
}

func Any(f func(string) bool, a []string) bool {
	for _, integer := range a {
		if f(integer) {
			return true
		}
	}
	return false
}

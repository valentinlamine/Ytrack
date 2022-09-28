package main

import "fmt"

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(isnumeric, tab1)
	answer2 := CountIf(isnumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
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

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	for _, integer := range tab {
		if f(integer) {
			result++
		}
	}
	return result
}

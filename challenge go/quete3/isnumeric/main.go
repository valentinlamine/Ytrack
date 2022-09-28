package main

import "fmt"

func main() {
	fmt.Println(isnumeric("HELLO"))
	fmt.Println(isnumeric("HELLO!"))
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

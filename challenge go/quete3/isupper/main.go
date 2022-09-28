package main

import "fmt"

func main() {
	fmt.Println(isupper("HELLO"))
	fmt.Println(isupper("HELLO!"))
}

func isupper(str string) bool {
	for index, letter := range str {
		_ = index //pour utiliser la valeur index
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}
	return true
}

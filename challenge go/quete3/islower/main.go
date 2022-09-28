package main

import "fmt"

func main() {
	fmt.Println(islower("hello"))
	fmt.Println(islower("hello!"))
}

func islower(str string) bool {
	for index, letter := range str {
		_ = index //pour utiliser la valeur index
		if letter < 'a' || letter > 'z' {
			return false
		}
	}
	return true
}

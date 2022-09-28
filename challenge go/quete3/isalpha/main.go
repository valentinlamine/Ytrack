package main

import "fmt"

func main() {
	fmt.Println(isalpha("Hello! How are you?"))
	fmt.Println(isalpha("HelloHowareyou"))
	fmt.Println(isalpha("What's this 4?"))
	fmt.Println(isalpha("Whatsthis4"))
}

func isalpha(str string) bool {
	for index, letter := range str {
		_ = index //pour utiliser la valeur index
		if (letter < 'A' || letter > 'Z') && (letter < 'a' || letter > 'z') && (letter < '0' || letter > '9') {
			return false
		}
	}
	return true
}

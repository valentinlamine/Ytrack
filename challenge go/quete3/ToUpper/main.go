package main

import "fmt"

func main() {
	fmt.Println(toupper("Hello! How are you ?"))
}

func toupper(str string) string {
	result := ""
	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			letter -= 32
		}
		result += string(letter)
	}
	return result
}

package main

import "fmt"

func main() {
	fmt.Println(tolower("Hello! How are you ?"))
}

func tolower(str string) string {
	result := ""
	for _, letter := range str {
		if letter >= 'A' && letter <= 'Z' {
			letter += 32
		}
		result += string(letter)
	}
	return result
}

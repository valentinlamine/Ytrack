// Write a function that capitalizes the first letter of each word and lowercases the rest.
// A word is a sequence of alphanumeric characters.package main
package main

import "fmt"

func main() {
	fmt.Println(capitalize("Hello! How are you? How+are+things+4you?"))
}

func capitalize(str string) string {
	result := ""
	for index, letter := range str {
		if index == 0 || !isalpha(string(str[index-1])) {
			if letter >= 'a' && letter <= 'z' {
				letter -= 32
			}
		} else if letter >= 'A' && letter <= 'Z' {
			letter += 32
		}
		result += string(letter)
	}
	return result
}

func isalpha(str string) bool {
	for _, letter := range str {
		if (letter < 'A' || letter > 'Z') && (letter < 'a' || letter > 'z') && (letter < '0' || letter > '9') {
			return false
		}
	}
	return true
}

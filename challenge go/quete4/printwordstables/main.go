package main

import "github.com/01-edu/z01"

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}

func SplitWhiteSpaces(str string) []string {
	result := []string{}
	word := ""
	for _, char := range str {
		if char == ' ' || char == '\t' || char == '\n' {
			result = append(result, word)
			word = ""
		} else {
			word += string(char)
		}
	}
	result = append(result, word)
	return result
}

func PrintWordsTables(a []string) {
	for _, strings := range a {
		for _, char := range strings {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

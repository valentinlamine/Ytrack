// 0 is returned when the first string equals the second string (a==b).
// 1 is returned when the first string is greater than the second string (a>b).
// -1 is returned when the first string is less than the second string (a<b).
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compare("Hello!", "Hello!"))
	fmt.Println(strings.Compare("Hello!", "Hello!"))
	fmt.Println(compare("Salut!", "lut!"))
	fmt.Println(strings.Compare("Salut!", "lut!"))
	fmt.Println(compare("Ola!", "Ol"))
	fmt.Println(strings.Compare("Ola!", "Ol"))
}

func compare(string1, string2 string) int {
	if string1 == string2 {
		return 0
	}
	for index, letter := range string1 {
		if index >= len(string2) {
			return 1
		}
		if letter > rune(string2[index]) {
			return 1
		}
		if letter < rune(string2[index]) {
			return -1
		}
	}
	return -1
}

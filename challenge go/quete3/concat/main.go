package main

import "fmt"

func main() {
	fmt.Println(concat("Hello", " World"))
}

func concat(str1, str2 string) string {
	result := ""
	for index, letter := range str1 + str2 {
		if index < len(str1+str2) {
			result = result + string(letter)
		}
	}
	return result
}

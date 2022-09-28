package main

import "fmt"

func main() {
	fmt.Println(lastrune("Hello World", 1))
}

func lastrune(str string, n int) rune {
	if n >= len(str) {
		return 0
	}
	for index, letter := range str {
		if index == n {
			return letter
		}
	}
	return 0
}

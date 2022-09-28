package main

import "fmt"

func main() {
	fmt.Println(lastrune("Hello World"))
}

func lastrune(str string) rune {
	for index, letter := range str {
		if index == len(str)-1 {
			return letter
		}
	}
	return 0
}

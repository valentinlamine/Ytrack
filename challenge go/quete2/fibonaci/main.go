package main

import "github.com/01-edu/z01"

func main() {
	resultat := fibonacci(3)
	z01.PrintRune(rune(resultat))
}

func fibonacci(number int) int {
	if number == 1 {
		return 1
	} else if number == 0 {
		return 0
	} else if number < 0 {
		return -1
	} else {
		return fibonacci(number-1) + fibonacci(number-2)
	}
}

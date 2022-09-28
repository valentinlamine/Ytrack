package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// Write a program that prints the corresponding letter in the n position of the latin alphabet, where n is each argument received.
// For example 1 matches a, 2 matches b, etc. If n does not match a valid position of the alphabet or if the argument is not an integer, the program should print a space (" ").
// A flag --upper should be implemented. When used, the program prints the result in upper case. The flag will always be the first argument
func main() {
	argms := os.Args
	upper := false
	for i := 1; i < len(argms); i++ {
		if argms[i] == "--upper" {
			upper = true
		} else {
			n, err := strconv.Atoi(argms[i])
			if err != nil {
				z01.PrintRune(' ')
			} else if n < 1 || n > 26 {
				z01.PrintRune(' ')
			} else {
				if upper {
					z01.PrintRune(rune(n + 64))
				} else {
					z01.PrintRune(rune(n + 96))
				}
			}
		}
	}
}

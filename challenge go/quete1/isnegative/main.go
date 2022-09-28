package main

import "github.com/01-edu/z01"

//func is negative
func main() {
	is_negative(1)
	is_negative(0)
	is_negative(-1)
}
func is_negative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}

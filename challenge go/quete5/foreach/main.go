package main

import "github.com/01-edu/z01"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(printnbr, a)

}

func printnbr(number int) {
	if number < 0 {
		z01.PrintRune('-')
		number = -number
	}
	if number > 9 {
		printnbr(number / 10)
	}
	z01.PrintRune(rune(number%10 + 48))
}

func ForEach(f func(int), a []int) {
	for _, integer := range a {
		f(integer)
	}
}

// Write a function that prints all possible combinations of n different digits in ascending order.
// n will be defined as : 0 < n < 10
//(for n = 1) '0, 1, 2, 3, ..., 8, 9'
// (for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'
//it must be a recusive function

package main

import "github.com/01-edu/z01"

func main() {
	printcombn(3)
}

func printcombn(n int) {
	if n > 0 && n < 10 {
		for i := 0; i < 10; i++ {
			printcomb(i, n-1)
		}
	}
}

func printcomb(n int, m int) {
	z01.PrintRune(rune(n + 48))
	if m > 0 {
		for i := n + 1; i < 10; i++ {
			printcomb(i, m-1)
		}
	} else {
		z01.PrintRune(10)
	}
}

package main

import "github.com/01-edu/z01"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for l := j + 1; l < 10; l++ {
					z01.PrintRune(rune(i + 48))
					z01.PrintRune(rune(j + 48))
					z01.PrintRune(32)
					z01.PrintRune(rune(k + 48))
					z01.PrintRune(rune(l + 48))
					if i != 9 || j != 8 || k != 9 || l != 9 {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}

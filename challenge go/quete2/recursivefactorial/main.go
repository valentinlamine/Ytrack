package main

import "fmt"

func main() {
	resultat := recursivefactorial(4)
	fmt.Println(resultat)
}

func recursivefactorial(number int) int {
	if number <= 1 {
		return 1
	} else {
		return number * recursivefactorial(number-1)
	}
}

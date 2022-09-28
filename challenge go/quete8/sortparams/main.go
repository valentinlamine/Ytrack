package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argms := os.Args
	//sort the arguments
	for i := 1; i < len(argms); i++ {
		for j := i + 1; j < len(argms); j++ {
			if argms[i] > argms[j] {
				argms[i], argms[j] = argms[j], argms[i]
			}
		}
	}
	//print the arguments
	for i := 1; i < len(argms); i++ {
		for _, c := range argms[i] {
			z01.PrintRune(c)
		}
		z01.PrintRune(10)
	}

}

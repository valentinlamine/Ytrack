package main

import "fmt"

func main() {
	fmt.Println(alphacount("Salut les amis moi je mange la glace"))
}

func alphacount(str string) int {
	counter := 0
	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			counter++
		}
	}
	return counter
}

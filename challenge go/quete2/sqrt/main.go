package main

import "fmt"

func main() {
	resultat2 := sqrt(19)
	fmt.Println(resultat2)
}
func sqrt(number int) int {
	for i := 1; i < number; i++ {
		if i*i == number {
			return i
		}
	}
	return 0
}

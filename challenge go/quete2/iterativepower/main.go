package main

import "fmt"

func main() {
	resultat := isiterativepower(5, 7)
	fmt.Println(resultat)
}

func isiterativepower(value1 int, value2 int) int {
	resultat := value1
	for i := 0; i < value2-1; i++ {
		resultat = resultat * value1
	}
	return resultat
}

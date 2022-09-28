package main

import "fmt"

func main() {
	fmt.Println(firstrune("Hello World"))
}

func firstrune(str string) rune {
	for indice, letter := range str { //on parcours le string
		if indice == 0 { //on souhaite récupérer le premier indice
			return letter //on le renvoie sous forme de rune
		}
	}
	return 0 //return qui ne sera jamais atteind
}

package main

import (
	"fmt"
)

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func UltimateDivMod(a *int, b *int) {
	temp1 := *a
	temp2 := *b
	*a = temp1 / temp2
	*b = temp1 % temp2
}

package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	count := 0
	for _, car := range s {
		_ = car
		count += 1
	}
	return count
}

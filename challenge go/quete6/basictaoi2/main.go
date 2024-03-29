package main

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
	var result int
	for _, each := range s {
		switch each {
		case '0':
			result = result*10 + 0
		case '1':
			result = result*10 + 1
		case '2':
			result = result*10 + 2
		case '3':
			result = result*10 + 3
		case '4':
			result = result*10 + 4
		case '5':
			result = result*10 + 5
		case '6':
			result = result*10 + 6
		case '7':
			result = result*10 + 7
		case '8':
			result = result*10 + 8
		case '9':
			result = result*10 + 9
		default:
			return 0
		}
	}
	return result
}

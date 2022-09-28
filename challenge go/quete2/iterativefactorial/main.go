package main

import "fmt"

func main() {
	result := iteractivefactorial(4)
	fmt.Println(result)

}

func iteractivefactorial(number int) int {
	result := number
	for i := number - 1; i > 0; i-- {
		result = result * i
	}
	return result
}

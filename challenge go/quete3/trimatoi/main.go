package main

import "fmt"

func main() {
	fmt.Println(trimatoi("12345"))
	fmt.Println(trimatoi("str123ing45"))
	fmt.Println(trimatoi("012 345"))
	fmt.Println(trimatoi("Hello World!"))
	fmt.Println(trimatoi("sd+x1fa2W3s4"))
	fmt.Println(trimatoi("sd-x1fa2W3s4"))
	fmt.Println(trimatoi("sdx1-fa2W3s4"))
	fmt.Println(trimatoi("sdx1+fa2W3s4"))
}

func trimatoi(str string) int {
	result := 0
	neg := false
	neg_block := false
	for _, letter := range str {
		if letter == '-' && neg_block == false {
			neg = true
		} else if letter >= '0' && letter <= '9' {
			result = result*10 + int(letter-'0')
			neg_block = true
		}
	}
	if neg {
		result = -result
	}
	return result
}

package main

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	reverse_s := ""
	for i := len(s) - 1; i > -1; i-- {
		reverse_s += string(s[i])
	}
	return reverse_s
}

package main

import "fmt"

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(concatparams(test))
}

func concatparams(args []string) string {
	result := ""
	for _, strings := range args {
		result += strings + "\n"
	}
	return result
}

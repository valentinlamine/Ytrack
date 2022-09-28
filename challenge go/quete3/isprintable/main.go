// Write a function that returns true if the string passed as a parameter contains only printable characters, otherwise, returns false.
package main

import "fmt"

func main() {
	fmt.Println(isprintable("Hello"))
	fmt.Println(isprintable("Hello\n"))

}

func isprintable(str string) bool {
	for _, v := range str {
		if v < 32 || v > 126 {
			return false
		}
	}
	return true
}

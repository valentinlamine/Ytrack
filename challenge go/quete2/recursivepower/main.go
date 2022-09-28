package main

import "fmt"

func main() {
	resultat := recusrivepower(4, 3)
	fmt.Println(resultat)
}

func recusrivepower(num1 int, num2 int) int {
	if num2 <= 1 {
		return num1
	} else {
		return num1 * recusrivepower(num1, num2-1)
	}
}

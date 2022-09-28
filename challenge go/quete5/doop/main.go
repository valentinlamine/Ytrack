package main

import (
	"os"
	"strconv"
)

func main() {
	first_number := os.Args[1]
	operator := os.Args[2]
	second_number := os.Args[3]
	//check if first_number & second_number are numbers
	if _, err := strconv.Atoi(first_number); err == nil {
		if _, err := strconv.Atoi(second_number); err == nil {
			//check if operator is valid
			if operator == "+" || operator == "-" || operator == "*" || operator == "/" || operator == "%" {
				//convert first_number & second_number to int
				first_number_int, _ := strconv.Atoi(first_number)
				second_number_int, _ := strconv.Atoi(second_number)
				//check if second_number is 0
				if second_number_int == 0 && operator == "/" || second_number_int == 0 && operator == "%" {
					println("No division by 0")
					return
				}
				//calculate
				switch operator {
				case "+":
					println(first_number_int + second_number_int)
				case "-":
					println(first_number_int - second_number_int)
				case "*":
					println(first_number_int * second_number_int)
				case "/":
					println(first_number_int / second_number_int)
				case "%":
					println(first_number_int % second_number_int)
				}
			} else {
				return
			}
		} else {
			return
		}
	} else {
		return
	}
}

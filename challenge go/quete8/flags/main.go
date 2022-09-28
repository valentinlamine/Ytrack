package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

/*
--insert

	-i
	       This flag inserts the string into the string passed as argument.

--order

	-o
	       This flag will behave like a boolean, if it is called it will order the argument.
*/
func main() {
	argms := os.Args[1:]
	Isinsert := false
	index_insert := 0
	Isorder := false
	str := ""
	if len(argms) == 0 {
		printhelp()
		return
	}
	for index, arg := range argms {
		if startwith(arg, "--insert") || startwith(arg, "-i") {
			Isinsert = true
			index_insert = index
		} else if startwith(arg, "--order") || startwith(arg, "-o") {
			Isorder = true
		} else if startwith(arg, "--help") || startwith(arg, "-h") {
			printhelp()
			return
		} else {
			str += arg
		}
	}
	if Isinsert {
		str = insert(str, argms[index_insert])
	} else if Isorder {
		str = order(str)
	}

	//print the arguments
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func insert(str, insert string) string {
	return str + insert[9:]
}

func printhelp() {
	fmt.Print("--insert\n	-i\n			This flag inserts the string into the string passed as argument.\n")
	fmt.Print("--order\n	-o\n			This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func order(str string) string {
	//sort the string
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] > str[j] {
				str = str[:i] + string(str[j]) + str[i+1:j] + string(str[i]) + str[j+1:]
			}
		}
	}
	return str
}

func startwith(str, start string) bool {
	//return true if str start with start
	if len(str) < len(start) {
		return false
	}
	if str[:len(start)] == start {
		return true
	}
	return false
}

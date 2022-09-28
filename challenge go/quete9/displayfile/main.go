package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(file))
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Too many arguments")
	}
}

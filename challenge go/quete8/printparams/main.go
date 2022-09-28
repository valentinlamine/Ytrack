package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		if i > 0 {
			fmt.Println(arg)
		}
	}
}

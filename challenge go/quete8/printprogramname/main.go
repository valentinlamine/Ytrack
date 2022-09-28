package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Program name:", filepath.Base(os.Args[0]))
}

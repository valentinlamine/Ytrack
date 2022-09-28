package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var ERRORS = map[string]string{
	"file_missing": "ERROR: %s: no such file or directory\n",
}

func error_handler(err interface{}, arg string) {
	fmt.Printf(ERRORS[err.(string)], arg)
}

func main() {
	args := os.Args[1:len(os.Args)]
	if len(args) == 0 {
		wait_input()
	}
	exec_args(args)
}

func wait_input() {
	for ok := true; ok; {
		var input string
		str, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println(str, err)
		}
		fmt.Println(input)
	}
}

func exec_args(args []string) {
	var cur_arg string = ""
	var arg_ptr *string = &cur_arg

	defer func() {
		if a := recover(); a != nil {
			error_handler(a, cur_arg)
		}
	}()

	for _, arg := range args {
		if _, err := os.Stat(arg); err != nil {
			*arg_ptr = arg
			panic("file_missing")
		}
		cat(arg)
	}
}

func cat(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	fmt.Println(string(content))
}

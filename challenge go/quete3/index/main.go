package main

import "fmt"

func main() {
	fmt.Println(index("salut les amis moi je mange la glace", "mange"))
}

func index(str, substr string) int {
	for index, letter := range str {
		if letter == rune(substr[0]) {
			for index2, letter2 := range substr {
				if letter2 != rune(str[index+index2]) {
					break
				} else if index2 == len(substr)-1 {
					return index
				}
			}
		}
	}
	return -1
}

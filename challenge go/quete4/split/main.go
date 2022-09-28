package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(str, sep string) []string {
	result := []string{}
	index := -1
	for i := 0; len(str) >= 1; i++ {
		index++
		if index == len(str) {
			result = append(result, str)
			break
		}
		if rune(str[index]) == rune(sep[0]) {
			if str[index:index+len(sep)] == sep {
				result = append(result, str[:index])
				str = str[index+len(sep):]
				index = -1 //remettre l'index à 0 pour la prochaine itération
			}
		}
	}
	return result
}

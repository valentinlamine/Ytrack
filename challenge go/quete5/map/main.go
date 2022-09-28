package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(Isprime, a)
	fmt.Println(result)
}

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, integer := range a {
		result = append(result, f(integer))
	}
	return result
}

func Isprime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

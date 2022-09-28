package isprime

func Isprime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func Findnextprime(number int) int {
	resultat := number + 1
	for i := number + 1; !Isprime(i); i++ {
		resultat = i + 1
	}
	return resultat
}

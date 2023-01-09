package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
** retourne une chaine de caractere en fonction de si l'int est divisible par 3, 5 ou les deux, sinon retourne l'int
 */
func fizzBuzz(number int) string {
	switch {
	case number%3 == 0 && number%5 == 0:
		return "FizzBuzz"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(number)
}

/*
** function qui vérfie si on donne un chiffre précis en entrée sinon on fait sur une fourchette de 1 à 100
 */
func checkIfArgs() {
	// si pas d'argument d'entré on fait de 1 à 100 le fizzBuzz
	if len(os.Args) < 2 {
		// sur une fourchette de 1 à 100
		for i := 1; i <= 100; i++ {
			fmt.Println(fizzBuzz(i))
		}
	} else {
		// sinon on traduit l'input entré
		num, _ := strconv.Atoi(os.Args[1])
		fmt.Println(fizzBuzz(num))
	}
}

func main() {
	checkIfArgs()
}

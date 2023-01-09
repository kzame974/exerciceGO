package main

import (
	"fmt"
	"strconv"
)

// imprime fizz buzz fizzbuzz ou le nombre en fonction des conditions
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

// soit on donne un chiffre précis en entrée soit on fait sur une fourchette de 1 à 100
func main() {
	// sur une fourchette de 1 à 100
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

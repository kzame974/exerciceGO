package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	val := 0

	for {
		fmt.Print("Enter number: ")
		_, err := fmt.Scanf("%d", &val)
		// gestion d'erreur
		if err != nil {
			panic(err)
			return
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recover following panic :", r)
			}
		}()
		// on exclu tous les types d'entré sauf les int entier sans le zéro
		verifyNumber := regexp.MustCompile("^[1-9][0-9]*$")

		// on vérifie la valeur entré par le user
		if verifyNumber.MatchString(strconv.Itoa(val)) {
			fmt.Println("You entered number:", val)
			break
		} else if val == 0 {
			fmt.Println("0 is neither negative nor positive")
		} else {
			panic("You entered wrong input")
		}
	}
}

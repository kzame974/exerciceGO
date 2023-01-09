package main

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
boucle qui permet de demander un int à l'utilisateur à chaque fois. Il redemande à chaque fois sauf si celui ci est négatif
alors on kill le programme avec panic et on récupère le message avec defer
*/
func askNumber() {
	val := 0

	for {
		fmt.Print("Enter number: ")
		_, err := fmt.Scanf("%d", &val)
		// gestion d'erreur sur le scan
		if err != nil {
			panic(err)
			return
		}
		// defer va récupérer une erreur jeter pour l'afficher en sortie
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recover following panic :", r)
			}
		}()
		verifyInputNumber(val)
	}
}

// on vérifie l'input rentré par le user
func verifyInputNumber(val int) {
	//on garde que les int entier zéro exclu
	verifyNumber := regexp.MustCompile("^[1-9][0-9]*$")

	// on fait un comparaison avec le regex imposé
	if verifyNumber.MatchString(strconv.Itoa(val)) {
		fmt.Println("You entered number:", val)
	} else if val == 0 {
		fmt.Println("0 is neither negative nor positive")
	} else {
		panic("You entered a negative number!")
	}
}

func main() {
	askNumber()
}

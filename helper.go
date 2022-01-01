package main

// para poder usar una packete que este fuera de main tenemos que declarar otro nombre del pacquete y importarlo conjuntamente desde el archivo de go.mod

import (
	"strings"
)

func validateUserInputs(firstName string, lastName string, email string, userTickets uint, remaningTickets uint) (bool, bool, bool) {
	// input validation
	// len() nos devuelve el numero de caracteres que tiene una cadena
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remaningTickets

	return isValidName, isValidEmail, isValidTickets
}

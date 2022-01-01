package main

import (
	"fmt"
	"sync"
	"time"
)

var confName = "Ticket go"

const confereceTicket = 50

var remaningTickets uint = 50
var bookingList = make([]UserData, 0)

// para poder tener objetos con datos mezaclados necesitamos una estructura
// debe de tener letra mayuscula
type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

// cuando creamos una funcion en segundo plano o en segundo hilo y el hilo principal no espera a que el hijo termine
// para eso usamos waitgroup que espera a que la rutina hijo termine su ejecucion
var wg = sync.WaitGroup{}

// wg ocupa 3 funciones, add, done, wait,
// Add con el numero de elementos que queremos que se ejecuten
// wait espera a que el hilo hijo termine que se pone al final de la funcion principal
// done le dice al hijo que ya termino su ejecucion y puede terminar y este va dentro de la funcion que lo llamo

func main() {

	// for {

	// saludo
	greetUsers()
	firstName, lastName, email, userTickets := getUserInputs()

	fmt.Println("#############################################")
	isValidName, isValidEmail, isValidTickets := validateUserInputs(firstName, lastName, email, userTickets, remaningTickets)

	if isValidName && isValidEmail && isValidTickets {

		// para mandar la funcion de ticket a otro hijo, usamos la palabra reservada go, lo que hace es crear un nuevo hilo
		// para que el hilo sea independiente del hilo principal de la aplicacion
		wg.Add(1)
		// Goroutine es una funcion que se ejecuta en segundo plano
		// el 1 es el numero de hilos que queremos que se ejecuten
		go sendTicket(userTickets, firstName, lastName, email)

		// actualizamos el numero de tickets disponibles
		remaningTickets = remaningTickets - userTickets

		fmt.Printf("Bienvenido %v %v, su correo es %v y tiene %v tickets\n", firstName, lastName, email, userTickets)
		fmt.Printf("Su pedido ha sido procesado, enviaramos sus tickets al correo %v\n", email)
		fmt.Printf("Quedan %v tickets disponibles\n", remaningTickets)

		fmt.Println("#############################################")

		name := printFristName()
		fmt.Printf("gracias por tu compra %v:\n", name)

		fmt.Println("#############################################")

		fmt.Printf("data: %v\n", bookingList)

	} else {
		fmt.Println("#############################################")
		fmt.Println("#                                           #")
		fmt.Println("#   ERROR: registre los valores correctos   #")
		fmt.Println("#                                           #")
		fmt.Println("#############################################")

		if !isValidName {
			fmt.Println("Por favor ingrese un nombre valido")
		}

		if !isValidEmail {
			fmt.Println("Por favor ingrese un correo valido")
		}

		if !isValidTickets {
			if userTickets > remaningTickets {
				fmt.Printf("Solo hay %v tickets disponibles\n", remaningTickets)

			} else {
				fmt.Println("Por favor ingrese un numero de tickets valido")
			}
		}
	}

	if remaningTickets == 0 {
		fmt.Println("Lo sentimos, no hay tickets disponibles")
		// break
		// para que el for termine
	}

	// se coloca el waitgroup para que el hilo principal espere a que el hilo hijo termine
	wg.Wait()

}

// }

func greetUsers() {
	// cuando pasamos un parametro a la funcion tenenmos que declarar el tipo de dato que tendra la variable
	fmt.Println("#############################################")
	fmt.Printf("#            %v                      #\n", confName)
	fmt.Println("#############################################")
	fmt.Println("#                                           #")
	fmt.Printf("%v En este sitio encontraras toda la informacion de la conferencia\n", confName)
	fmt.Println("#                                           #")
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Por favor ingrese su nombre")
	fmt.Scanf("%v", &firstName)

	fmt.Println("Por favor ingrese su apellido")
	fmt.Scanf("%v", &lastName)

	fmt.Println("Por favor ingrese su correo")
	fmt.Scanf("%v", &email)

	fmt.Println("Por favor ingrese el numero de tickets")
	fmt.Scanf("%v", &userTickets)

	// var userData = make(map[string]string)
	// seteamos los valores que queremos que tenga el mapa
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	userData.firstName = firstName
	userData.lastName = lastName
	userData.email = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// al cambiar el userdata por un struct tenemos que usar el nombre de la variable como un objeto de javascript
	userData.userTickets = userTickets

	bookingList = append(bookingList, userData)

	return firstName, lastName, email, userTickets
}

func printFristName() []string {
	// para poder retornar un valor de la funcion tenemos que declarar su tipo de dato

	firstName := []string{}
	for _, bookingList := range bookingList {
		// para ignorar una variable que no ocuparemos usamos el guion bajo
		firstName = append(firstName, bookingList.firstName)
	}

	return firstName
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	// el paquete time con el parametro time.Second emula un setTimeout de javascript
	// time.Sleep para el trabajo en hilo de la aplicacion
	time.Sleep(10 * time.Second)
	// con Sprintf podemos crar una variable y a su vez imprimir en consola
	var ticket = fmt.Sprintf(" %v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#############################################")
	fmt.Printf("Sendig ticket:\n %v \nto email addres %v...\n", ticket, email)
	fmt.Println("#############################################")
	wg.Done()
}

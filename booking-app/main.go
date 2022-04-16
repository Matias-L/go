// Inicio el proyecto con el comando go mod init booking-app
package main

import "fmt"

func main() {
	var conferenceName string = "Most bestest conference"
	const maxTickets = 50
	var ticketsRestantes uint = 50

	fmt.Println("Hola guacho. Bienvenido a booking app")
	fmt.Printf("Aca podes reservar tickets para %v re piolamente\n", conferenceName)
	fmt.Printf("Para que sepas, de los %v tickets quedan %v\n", maxTickets, ticketsRestantes)

	var userName string
	var userTickets int = 2

	fmt.Scan(&userName)

	fmt.Printf("Usuario %v tiene %v tickets.\n", userName, userTickets)
}

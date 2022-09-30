package main

import (
	"fmt"
	"os"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

var FILE_PATH = getFilePath()

func getFilePath() string {
	workingDirectory, _ := os.Getwd()
	return fmt.Sprintf("%s/%s", workingDirectory, "tickets.csv")
}

func main() {
	var tickets []service.Ticket

	file := file.File{Path: FILE_PATH}

	tickets, err := file.Read()

	if err != nil {
		fmt.Println("No se pudo leer el archivo")
		return
	}

	// Funcion para obtener tickets del archivo csv
	bookingsServe := service.NewBookings(tickets)

	ticket, err := bookingsServe.Read(10)

	fmt.Printf("%+v | %+v\n", ticket, err)

	newTicket := service.Ticket{
		Names: "Daniel Medina",
		Id:    20,
	}

	bookingsServe.Create(newTicket)
}

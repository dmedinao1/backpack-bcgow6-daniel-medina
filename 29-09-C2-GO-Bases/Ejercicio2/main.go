package main

import (
	"fmt"

	"github.com/dmedinao1/backpack-bcgow6-daniel-medina/customers/services"
)

func main() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}

		fmt.Println("No han quedado archivos abiertos")
		fmt.Println("Fin de la ejecución")
	}()

	var FullName string
	var DNI string
	var PhoneNumber string
	var Adress string

	fmt.Print("Digita el nombre: ")
	fmt.Scan(&FullName)

	fmt.Print("Digita el DNI: ")
	fmt.Scan(&DNI)

	fmt.Print("Digita el telefono: ")
	fmt.Scan(&PhoneNumber)

	fmt.Print("Digita la dirección: ")
	fmt.Scan(&Adress)

	savedCustomer, err := services.SaveNewClient(FullName, DNI, PhoneNumber, Adress)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Cliente guardado: %+v\n", savedCustomer)

}

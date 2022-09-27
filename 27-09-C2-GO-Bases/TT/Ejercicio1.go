package main

import "fmt"

/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura.
Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.

*/

type User struct {
	Name     string
	Lastname string
	Age      int
	Email    string
	Pass     string
}

func (user User) changeName(newName, newLastname string) {
	fmt.Printf("Dirección en changeName: %p\n", &user)
	user.Name = newName
	user.Lastname = newLastname
}

func main() {
	myUser := User{Name: "Daniel", Lastname: "Medina"}
	fmt.Printf("%+v\n", myUser)
	fmt.Printf("Dirección en main antes de changeName: %p\n", &myUser)

	myUser.changeName("Miguel", "Ortega")

	fmt.Printf("%+v\n", myUser)
	fmt.Printf("Dirección en main despues de changeName: %p\n", &myUser)
}

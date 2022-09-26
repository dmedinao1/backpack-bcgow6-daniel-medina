package main

import (
	"fmt"
	"time"
)

/*
Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

*/

type Student struct {
	Name      string
	Lastnames string
	DNI       int
	BirthDay  time.Time
}

func (student Student) detail() {
	fmt.Printf("Nombre: %v\n", student.Name)
	fmt.Printf("Apellido: %v\n", student.Lastnames)
	fmt.Printf("DNI: %v\n", student.DNI)
	fmt.Printf("Fecha: %v\n", student.BirthDay)
}

func main() {
	myStudent := Student{Name: "Daniel", BirthDay: time.Date(2001, 3, 19, 14, 30, 45, 100, time.Local)}
	myStudent.detail()
}

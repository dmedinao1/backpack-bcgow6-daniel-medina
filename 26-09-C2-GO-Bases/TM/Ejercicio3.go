package main

import (
	"fmt"
	"strings"
)

/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

var (
	bonusByCategory = map[string]float64{
		"b": .2,
		"a": .5,
	}
	salariesByCategory = map[string]float64{"c": 1_000, "b": 1_500, "a": 3_000}
)

func main() {
	var workedMinutes int

	fmt.Print("Digite los minutos trabajados por el trabajador: ")
	fmt.Scan(&workedMinutes)

	var category string

	fmt.Print("Digite la categoria del trabajador[a,b,c]: ")
	fmt.Scan(&category)

	categoryLower := strings.ToLower(category)

	if categoryLower != "a" && categoryLower != "b" && categoryLower != "c" {
		fmt.Printf("Categoría %v no es válida\n", category)
		return
	}

	fmt.Printf("Salario del trabajador es: %v", getEmployeeSalary(workedMinutes, categoryLower))

}

func getEmployeeSalary(workedMinutes int, category string) float64 {
	hours := getHoursFromMinutes(workedMinutes)

	bonus := bonusByCategory[category]

	baseSalaryTotal := hours * salariesByCategory[category]

	return baseSalaryTotal + baseSalaryTotal*bonus
}

func getHoursFromMinutes(workedMinutes int) float64 {
	return float64(workedMinutes) / 60
}

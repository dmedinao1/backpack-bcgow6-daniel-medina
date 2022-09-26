package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/
func main() {
	fmt.Println(calculateAvg())
}

func calculateAvg(studentScores ...float64) (float64, error) {
	if len(studentScores) == 0 {
		return 0, nil
	}

	if hasANegativeNumber(studentScores) {
		return 0, errors.New("La operación no se pudo realizar debido a que fue ingresado un número negativo")
	}

	var sum float64

	for _, score := range studentScores {
		sum += score
	}

	return sum / float64(len(studentScores)), nil

}

func hasANegativeNumber(studentScores []float64) bool {
	for _, score := range studentScores {
		if score < 0 {
			return true
		}
	}

	return false
}

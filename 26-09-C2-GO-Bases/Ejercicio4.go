package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

*/

func main() {
	minFunc, _ := operation(minimum)
	//averageFunc, _ := operation(average)
	//maxFunc, _ := operation(maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5, -1)
	//averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	//maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("MinValue: %v\n", minValue)

}

func operation(operation string) (func(values ...float64) float64, error) {
	switch operation {
	case minimum:
		return func(values ...float64) float64 {
			var minimum float64

			for i, value := range values {
				if i == 0 {
					minimum = value
					continue
				}

				if value < minimum {
					minimum = value
				}
			}

			return minimum
		}, nil
	case average:
		return func(values ...float64) float64 {
			if len(values) == 0 {
				return 0
			}

			var sum float64

			for _, score := range values {
				sum += score
			}

			return sum / float64(len(values))

		}, nil

	case maximum:
		return func(values ...float64) float64 {
			var minimum float64

			for i, value := range values {
				if i == 0 {
					minimum = value
					continue
				}

				if value > minimum {
					minimum = value
				}
			}

			return minimum
		}, nil
	default:
		return nil, errors.New("Función para la operación no encontrada")
	}

}

func hasANegativeNumber(studentScores []float64) bool {
	for _, score := range studentScores {
		if score < 0 {
			return true
		}
	}

	return false
}

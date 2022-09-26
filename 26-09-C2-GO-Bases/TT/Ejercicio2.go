package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	Values [][]float64
	Height int
	Width  int
}

func (m Matrix) Init(height, widht int) {
	m.Height = height
	m.Width = widht

	for i := 0; i < height; i++ {
		m.Values = append(m.Values, make([]float64, widht))
	}
}

func (m Matrix) GetMaxValue() float64 {
	var maxValue float64
	for i, row := range m.Values {
		for j, item := range row {
			if i == 0 && j == 0 {
				maxValue = item
			}

			if item > maxValue {
				maxValue = item
			}
		}
	}
	return maxValue
}

func (m Matrix) Print() {
	for _, row := range m.Values {
		fmt.Println(row)
	}
}

func (m Matrix) Set(values ...float64) error {
	if matrixCapacity := m.Height * m.Width; len(values) > matrixCapacity {
		return errors.New("La matriz no tiene capacidad suficiente")
	}

	var index int

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			if index >= len(values) {
				m.Height = 1
				return nil
			}

			m.Values[i][j] = values[index]
			index++
		}
	}

	return nil
}

func main() {
	myMatrix := Matrix{}
	myMatrix.Init(3, 3)
	err := myMatrix.Set(1, 2, 3, 4, 5, 6)

	if err != nil {
		fmt.Println(err)
		return
	}

	myMatrix.Print()
}

package main

import "fmt"

/*
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
	-Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
	-Luego imprimí cada una de las letras.
*/

func main() {
	var palabra string
	fmt.Print("Digita una palabra: ")
	fmt.Scan(&palabra)

	fmt.Printf("La palabra '%v' tiene %v letras\n", palabra, len(palabra))

	for _, letra := range palabra {
		fmt.Printf("%v\n", string(letra))
	}
}

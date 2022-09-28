package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

func main() {
	var salary int

	fmt.Print("Ingrese el salario: ")
	_, err := fmt.Scan(&salary)

	if err != nil {
		fmt.Println("No se pudo leer la entrada")
		return
	}

	err = validateSalary(salary)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Debe pagar impuesto")

}

func validateSalary(salary int) error {
	if salary < 150_000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}

	return nil
}

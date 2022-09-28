package main

import "fmt"

/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

type SalaryError struct {
}

func (SalaryError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

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
		return SalaryError{}
	}

	return nil
}

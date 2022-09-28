package main

import "fmt"

/*
Ejercicio 3 - Impuestos de salario #3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”,
para que el mensaje de error reciba por parámetro el valor de “salary”
indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir:
	“error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
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
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
	}

	return nil
}

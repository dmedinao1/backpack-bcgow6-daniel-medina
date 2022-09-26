package main

import "fmt"

/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
*/
func main() {
	var salary float64

	fmt.Print("Ingrese el salario del trabajador: ")
	_, err := fmt.Scan(&salary)

	if err != nil {
		fmt.Println("No se pudo leer la entrada")
		return
	}

	fmt.Printf("El impuesto a aplicar al trabajador que gana %v es: %.2f\n", salary, getTastBySalary(salary))

}

func getTastBySalary(salary float64) float64 {
	var discountToApply float64

	if salary > 50_000 {
		discountToApply = 0.17
	}

	if salary > 150_000 {
		discountToApply += 0.1
	}

	return salary * float64(discountToApply)
}

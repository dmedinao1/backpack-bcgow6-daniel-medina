package main

import (
	"errors"
	"fmt"
)

/*
Bonus Track -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo.
Desarrolla las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
Calcular el medio aguinaldo correspondiente al trabajador
Fórmula de cálculo de aguinaldo:
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.
*/

type EmployeeSalaryData struct {
	workedHours           int
	hourValue             float64
	bestSalaryInSemester  float64
	workedHoursInSemester int
}

func (e EmployeeSalaryData) calculateSalary() (float64, error) {
	var calculatedSalary float64

	if e.workedHours < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	if e.hourValue < 0 {
		return 0, fmt.Errorf("error: el valor de la hora del %v trabajador no es válida", e.hourValue)
	}

	calculatedSalary = float64(e.workedHours) * e.hourValue

	if calculatedSalary >= 150_000 {
		calculatedSalary -= calculatedSalary * 0.1
	}

	return calculatedSalary, nil
}

func (e EmployeeSalaryData) calculateBonus() (float64, error) {
	if e.bestSalaryInSemester < 0 || e.workedHoursInSemester < 0 {
		return 0, errors.New("error: el mejor salario del trabajador y las horas trabajadas en el semestre deben ser mayores a cero")
	}

	return e.bestSalaryInSemester / 12 * float64(e.workedHoursInSemester), nil
}

func main() {

	myEmployee := EmployeeSalaryData{workedHours: 120, hourValue: 100}

	salary, errSalary := myEmployee.calculateSalary()

	if errSalary != nil {
		fmt.Println(errSalary)
	}

	fmt.Printf("Salario del trabajador: %v\n", salary)

	bonus, errBonus := myEmployee.calculateBonus()

	if errBonus != nil {
		fmt.Println(errSalary)
	}

	fmt.Printf("Medio aguinaldo del trabajador: %v\n", bonus)
}

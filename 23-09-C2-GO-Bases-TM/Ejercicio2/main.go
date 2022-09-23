package main

import "fmt"

/*
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.

Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/
func main() {

	fmt.Println("\t*VERIFICACIÓN DE CRÉDITOS*")
	fmt.Println("")

	// Leyendo edad del cliente
	var edadCliente int
	fmt.Print("Ingrese la edad del client: ")
	_, err := fmt.Scan(&edadCliente)

	if err != nil {
		fmt.Println("No se pudo leer la entrada")
		return
	}

	if edadCliente < 22 {
		fmt.Println("La edad requerida es de 22 años")
		return
	}

	fmt.Println("")

	// Leyendo estado del cliente
	var estadoClienteEntrada string
	fmt.Print("¿El cliente es empleado activo? [s/n] ")
	_, errEstado := fmt.Scan(&estadoClienteEntrada)

	if errEstado != nil {
		fmt.Println("No se pudo leer la entrada")
		return
	}

	if estadoClienteEntrada != "s" {
		fmt.Println("El cliente debe ser un empleado")
		return
	}

	fmt.Println("")

	// Leyendo antiguedad del cliente
	var antiguedadCliente int
	fmt.Print("Ingrese la antiguedad del cliente: ")
	_, errAntiguedad := fmt.Scan(&antiguedadCliente)

	if errAntiguedad != nil {
		fmt.Println("No se pudo leer la entrada")
		return
	}

	if antiguedadCliente < 1 {
		fmt.Println("La antiguedad requerida es de 1 año")
		return
	}

	fmt.Println("")

	// Leyendo sueldo del cliente
	var sueldoCliente int
	fmt.Print("Ingrese el sueldo del cliente: ")
	_, errSueldo := fmt.Scan(&sueldoCliente)

	if errSueldo != nil {
		fmt.Println("No se pudo leer la entrada")
		return
	}

	if sueldoCliente > 100_000 {
		fmt.Println("Al cliente NO se le cobraran intereses")
	} else {
		fmt.Println("Al cliente se le cobraran intereses")
	}

	fmt.Println("")

	fmt.Println("Al cliente se le puede ofrecer un préstamo")
}

package main

import "fmt"

/*
Realizar una aplicación que contenga una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
*/
func main() {
	meses := [13]string{"", "Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	var mesConsulta int

	fmt.Print("Digita el número del mes a consultar [1-12]: ")

	_, error := fmt.Scan(&mesConsulta)

	if error != nil {
		fmt.Println("No se pudo leer la entrada")
		return
	}

	if mesConsulta < 1 || mesConsulta > 12 {
		fmt.Println("Entrada no reconocida")
		return
	}

	fmt.Printf("El mes que corresponde al número %d, es %v\ns", mesConsulta, meses[mesConsulta])
}

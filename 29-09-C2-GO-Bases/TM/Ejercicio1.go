package main

import (
	"fmt"
	"os"
)

/*

Ejercicio 1 - Datos de clientes
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo .txt.
Es necesario desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo, no han pasado el archivo a leer por nuestro programa.
Desarrolla el código necesario para leer los datos del archivo llamado “customers.txt” (recuerda lo visto sobre el pkg “os”).
Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al intentar leer un archivo que no existe, mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/

var FILE_PATH = getFilePath()

func getFilePath() string {
	workingDirectory, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s", workingDirectory, "myFiles", "customers.txt")
}

func main() {
	myFile, err := os.ReadFile(FILE_PATH)

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	fmt.Printf("Archivo leido con la longitud de %d bytes\n", len(myFile))
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados,
con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad),
el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)
*/

var FILE_PATH = getFilePath()

func getFilePath() string {
	workingDirectory, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s", workingDirectory, "myFiles", "products.csv")
}

func printHead(head string) {
	tokens := strings.Split(head, ",")
	fmt.Printf("%s\t%s\t%s\n", tokens[0], tokens[1], tokens[2])
}

func printRow(row string) {
	tokens := strings.Split(row, ",")
	price, err := strconv.ParseFloat(tokens[1], 64)

	if err != nil {
		return
	}

	fmt.Printf("%v\t%.2f\t%v\n", tokens[0], price, tokens[2])
}

func main() {
	data, err := os.ReadFile(FILE_PATH)

	if err != nil {
		fmt.Println(err)
	}

	rows := strings.Split(string(data), "\n")

	if len(rows) < 2 {
		fmt.Println("Archivo vacío o sin contenido suficiente")
		return
	}

	printHead(rows[0])

	for _, product := range rows[1:] {
		if product != "" {
			printRow(product)
		}
	}

}

package main

import (
	"fmt"
	"os"
)

/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.

*/

var FILE_PATH = getFilePath()

func getFilePath() string {
	workingDirectory, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s", workingDirectory, "myFiles", "products.csv")
}

type Product struct {
	id       int
	price    float64
	quantity int
}

func (p Product) toCSVRow() string {
	return fmt.Sprintf("%v,%.2f,%v\n", p.id, p.price, p.quantity)
}

func getCSVHeader() string {
	return "Id,Precio,Cantidad\n"
}

func getProudcts() []Product {
	return []Product{
		{id: 1, price: 10_000, quantity: 5},
		{id: 2, price: 50_000, quantity: 21},
		{id: 3, price: 101_000, quantity: 8},
		{id: 4, price: 210_000, quantity: 1},
	}
}

func main() {
	var productsData = []byte{}

	productsData = append(productsData, []byte(getCSVHeader())...)

	for _, product := range getProudcts() {
		productsData = append(productsData, []byte(product.toCSVRow())...)
	}

	err := os.WriteFile(FILE_PATH, productsData, 0644)

	if err != nil {
		fmt.Println(err)
	}

}

// func main() {
// 	os.Remove(FILE_PATH)

// 	file, err := os.Open(FILE_PATH)

// 	if err != nil {
// 		file, err = os.Create(FILE_PATH)

// 		if err != nil {
// 			panic("No se pudo crear el archivo")
// 		}

// 	}

// 	writer := bufio.NewWriter(file)
// 	writer.Write([]byte(getCSVHeader()))

// 	for _, product := range getProudcts() {
// 		writer.Write([]byte(product.toCSVRow()))
// 	}

// 	writer.Flush()
// 	file.Close()
// }

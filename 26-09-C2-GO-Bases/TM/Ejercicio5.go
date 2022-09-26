package main

import "fmt"

/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

*/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

var foodQuantityByAnimal = map[string]int{
	"dog":       10_000,
	"cat":       5_000,
	"hamster":   250,
	"tarantula": 150,
}

func main() {
	animalDog, msg := Animal("not")

	if msg != "" {
		fmt.Println(msg)
		return
	}

	var amount float64
	amount += animalDog(5)

	fmt.Println(amount)
}

func Animal(animal string) (func(int) float64, string) {
	foodWeightRequired, ok := foodQuantityByAnimal[animal]

	if !ok {
		return nil, "Animal no encontrado"
	}

	return buildAnimalFoodFunction(foodWeightRequired), ""
}

func buildAnimalFoodFunction(quantity int) func(int) float64 {
	return func(numOfAnimals int) float64 {
		return float64(numOfAnimals * quantity)
	}
}

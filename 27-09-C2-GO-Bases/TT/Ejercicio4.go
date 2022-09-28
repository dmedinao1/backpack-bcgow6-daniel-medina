package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 Ejercicio 4 - Ordenamiento
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand
package main

import (
   "math/rand"
)


func main() {
   variable1 := rand.Perm(100)
   variable2 := rand.Perm(1000)
   variable3 := rand.Perm(10000)
}

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo


*/

/*Constantes*/
const (
	Insertion = "Insertion"
	Bubble    = "Bubble"
	Selection = "Selection"
)

/*Algoritmos*/
func InsertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

func BubbleSort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}

func SelectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

/*END Algoritmos*/

/*structs*/
type AlgorithmsTable struct {
	InsertionTime int64
	BubbleTime    int64
	SelectionTime int64
}

/*END structs*/

func ExecuteAlgorithm(name string, channel chan string, items []int, sortFunction func([]int)) {
	sortFunction(items)
	channel <- name
}

func ExecuteAlgorithms(itemsToSort []int) AlgorithmsTable {
	channel := make(chan string)
	startTimeByAlgorithms := map[string]time.Time{}
	durationByAlgorithms := map[string]int64{}

	startTimeByAlgorithms[Bubble] = time.Now()
	go ExecuteAlgorithm(Bubble, channel, itemsToSort, BubbleSort)

	startTimeByAlgorithms[Insertion] = time.Now()
	go ExecuteAlgorithm(Insertion, channel, itemsToSort, InsertionSort)

	startTimeByAlgorithms[Selection] = time.Now()
	go ExecuteAlgorithm(Selection, channel, itemsToSort, SelectionSort)

	for i := 0; i < 3; i++ {
		terminatedAlgorithm := <-channel
		algorithmStartTime := startTimeByAlgorithms[terminatedAlgorithm]
		durationByAlgorithms[terminatedAlgorithm] = time.Now().Sub(algorithmStartTime).Milliseconds()
	}

	return AlgorithmsTable{
		InsertionTime: durationByAlgorithms[Insertion],
		BubbleTime:    durationByAlgorithms[Bubble],
		SelectionTime: durationByAlgorithms[Selection],
	}
}

func main() {
	for _, itemsCount := range []int{100_000} {
		fmt.Printf("Duración en Milisegundos de los algoritmos con %v elementos: %+v\n\n", itemsCount, ExecuteAlgorithms(rand.Perm(itemsCount)))
	}
}

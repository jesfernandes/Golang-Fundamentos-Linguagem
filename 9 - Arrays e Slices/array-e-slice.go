package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	// Para utilizar o array em Go,
	// o tamanho da lista precisa ser especificado
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// a principal diferença de um slice para o array em Go,
	// é que o tamanho da lista nao precisa ser especificado
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// outro fato é que podemos incluir um array em um slice
	// pois o slice é composto de arrays internos
	// para isso utilizamos o "append"
	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)
}

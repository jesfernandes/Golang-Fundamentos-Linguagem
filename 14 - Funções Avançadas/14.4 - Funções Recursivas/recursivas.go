package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {

	fmt.Println("Funções Recursivas")
	//Sequencia fibonacci para o exercicio
	// 1 1 2 3 5 8 13
	//Deve retornar um numero de acordo com a posição informada

	position := uint(12)
	fmt.Println((fibonacci(position)))

	fmt.Println("Imprimindo a sequencia de Fibonacci até a Posição informada")

	for i := uint(1); i <= position; i++ {
		fmt.Println(fibonacci(i))
	}

}

package main

import (
	"fmt"
)

//Novidades do Go , hoje os funções podem retornar parametros com tipos diferentes
func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculateMatematic(n1, n2 int8) (int8, int8) {
	somaTest := n1 + n2
	subtracaoTest := n1 - n2
	return somaTest, subtracaoTest
}

func main() {
	result := soma(10, 20)
	fmt.Println(result)

	var callFunction = func(text string) string {
		fmt.Println(text)
		return text

	}

	resultFunction := callFunction("texto inserido na função callFunction")
	fmt.Println(resultFunction)

	resultadoSoma, resultadoSubtracao := calculateMatematic(10, 20)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Exemplo de como nao usar um dos retornos de uma função
	// resultadoSoma, _ := calculateMatematic(10, 20)
	// fmt.Println(resultadoSoma)
}

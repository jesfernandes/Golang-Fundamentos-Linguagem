package main

import (
	"fmt"
)

func main() {

	fmt.Println("Estruturas de Controle")

	n1 := 1

	//ao contrario do java, nao precisamos usar parentes
	//a menos que queira executar em uma determinada ordem as condições , caso exista mais de uma..
	if n1 >= 15 {
		fmt.Println("Numero1 maior ou igual a 15")
	} else {
		fmt.Println("Numero1 menor que 15")
	}

	//if init
	//a variável só existe dentro do escopo (if)
	if n2 := n1; n2 > 0 {
		fmt.Println("Numero2 é maior do que 0")
	} else {
		fmt.Println("Numero2 é menor que 0")
	}
	// o else if também pode ser usado, dependendo da sua estrutura

}

package main

import "fmt"

func main() {

	var variavel1 int = 1
	var variavel2 int = variavel1

	fmt.Println("v1 ", variavel1, "v2 ", variavel2)

	variavel2++
	fmt.Println("v1 ", variavel1, " v2 com incremento: ", variavel2)

	// ponteiro é uma referencia de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 10
	ponteiro = &variavel3

	fmt.Println("v3 ", variavel3, " ponteiro: ", ponteiro)

	variavel3 = 25
	fmt.Println("v3 ", variavel3, " ponteiro: ", *ponteiro) //desreferenciação
}

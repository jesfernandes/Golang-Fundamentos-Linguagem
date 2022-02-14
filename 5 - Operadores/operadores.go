package main

import "fmt"

func main() {
	// Operadores aritméticos
	fmt.Println("==== Operadores aritméticos ====")

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	// restoDaDivisao := 10 \ 2
	mutiplicacao := 10 * 2

	fmt.Println(soma, subtracao, divisao, mutiplicacao)

	var n1 int16 = 10
	// var n2 int32 = 20
	var n3 int16 = 25
	//nao podemos calcular dados com tipos diferentes
	// somaNumeros := n1 + n2

	result := n1 + n3
	fmt.Println(result)

	// Fim dos operadores Aritméticos

	// Atribuição
	fmt.Println("==== Atribuição ====")
	var v1 string = "String"
	v2 := "String"
	fmt.Println(v1, v2)
	// Fim da Atribuição

	// Operadores Relacionais
	fmt.Println("==== Operadores Relacionais ====")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// Fim dos operadores relacionais

	// Operadores lógicos
	fmt.Println("==== Operadores Lógicos ====")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// Fim dos operadores lógicos

	// Operadores Unários
	fmt.Println("==== Operadores Unários ====")
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)

	numero--
	fmt.Println(numero)
	numero -= 5
	fmt.Println(numero)

	numero *= 5
	fmt.Println(numero)
	numero -= 5
	fmt.Println(numero)

	numero /= 5
	fmt.Println(numero)
	numero %= 5
	fmt.Println(numero)

	// Fim dos Operadores Unários

	// Operadores ternários nao existem em Go
}

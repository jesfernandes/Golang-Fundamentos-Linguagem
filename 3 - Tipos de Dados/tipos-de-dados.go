package main

import (
	"errors"
	"fmt"
)

func main() {

	//numeros inteiros, 4 tipos em Go - Quantidade de bits que eles ocupam
	//int8, int16, int32, int64

	var numero int32 = 100
	fmt.Println(numero)

	numero1 := 1000000000000
	fmt.Println(numero1)

	numero2 := -1000000000000
	fmt.Println(numero2)

	// uint32 nao aceita negativos  ex: -10000 ou qualquer outro uint
	var numero3 uint32 = 10000
	fmt.Println(numero3)

	// alias para uint32,
	var numero4 rune = 12345
	fmt.Println(numero4)

	// alias para uint8,
	var numero5 byte = 123
	fmt.Println(numero5)

	//reais
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	//inferencia de tipo para real
	numeroReal3 := 12345.56
	fmt.Println(numeroReal3)

	//Para Go nao existe o caracter char
	//mas com as apas simples é possivel referenciar ao que seria o char

	textoStr := "Texto"
	fmt.Sprintln(textoStr)

	character := 'a'
	fmt.Sprintln(character)

	//valor zero, variavel nao inicializada
	var textoNulo string
	fmt.Println(textoNulo)

	//boolean
	var boolean bool
	fmt.Println(boolean)

	//Usa-se um pacote nativo do Go chamado Errors
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}

package main

import "fmt"

func main() {

	var variavel1 string = "Variável String / Declaração Explícita "
	fmt.Println(variavel1)

	variavel2 := "Variável String com tipo determinado pelo valor / Declaração implicita"
	fmt.Println(variavel2)

	var (
		variavel3 string = "v3 "
		variavel4 string = "v4 "
	)

	fmt.Println(variavel3, variavel4+" sem inferencia de dados")

	variavel5, variavel6 := " v5", " v6"

	fmt.Println("com Inferência de dados, imprimindo "+variavel5, variavel6)

	const constante1 string = "CONSTANTE 1"
	fmt.Println(constante1 + " imprimindo constante")

	variavel6, variavel5 = variavel5, variavel6
	fmt.Println("substituindo valores das variaveis v6 e v5, imprimindo "+variavel6, variavel5)

}

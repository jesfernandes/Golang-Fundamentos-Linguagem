package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá mundo !")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando parametro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido como Parâmetro o valor -> %s", texto)
	}("string teste")

	fmt.Println(retorno)
}

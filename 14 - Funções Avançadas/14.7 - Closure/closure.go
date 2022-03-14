package main

import "fmt"

func closure() func() {
	texto := "In function closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	fmt.Println("Closure")
	texto := "Value initial"
	fmt.Println(texto)

	newFunction := closure()
	newFunction()
}

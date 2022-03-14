package main

import "fmt"

//funcao init pode ser uma por arquivo e nao como a main que é uma por pacote,
//nao importa a orderm dos metodos, a funcao init sempre sera executada primeiro
func init() {
	fmt.Println("Executing function initial")
}

func main() {
	fmt.Println("Executing function main")
}

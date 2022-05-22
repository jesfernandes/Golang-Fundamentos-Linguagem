package main

import (
	"fmt"
	"time"
)

func main() {
	//Concorrência != Paralelismo
	go escrever("Olá mundo!") //goroutine, executa essa função junto com as outras funções simultaneamente
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

// Pode-se ter várias 1 - goroutines
// Para implementar, deve ser usado o comando "go" antes do metódo a ser executado.

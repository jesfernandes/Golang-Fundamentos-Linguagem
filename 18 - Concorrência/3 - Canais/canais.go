package main

import (
	"fmt"
	"time"
)

func main() {

	//um canal pode receber informacoes
	//um canal pode passar informacoes

	canal := make(chan string)

	go escrever("Olá mundo! ", canal)

	fmt.Println("Depois da funcao escrever comecar a ser executada")
	for mensagem := range canal { //ele espera que um valor seja enviado para o canal a partir de uma variavel
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //atribuindo a um canal um valor de uma variavel
		time.Sleep(time.Second)
	}

	close(canal)
}

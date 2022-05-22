package main

import "fmt"

func main() {

	canal := make(chan string, 2) // canal com buffer, só bloqueia, quando atingir a capacidade máxima
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	//canal <- "Terceiro valor!" //aqui ele gera o deadlock

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}

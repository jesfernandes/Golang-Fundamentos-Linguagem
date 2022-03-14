package main

import "fmt"

func funcao1() {
	fmt.Println("Execute first function example")
}

func funcao2() {
	fmt.Println("Execute second function example")
}

func mediaStudent(n1, n2 float32) bool {
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func verifyStudentApproval(n1, n2 float32) bool {
	defer fmt.Println("Complete average calculation.")
	fmt.Println("Entering the function to check if the student is approved:")
	return mediaStudent(n1, n2)
}

func main() {

	fmt.Println("Defer (Adiar)")
	//principais usos de uma instrução defer é o da limpeza de recursos,
	//como arquivos abertos, conexões de rede e conexões de banco de dados
	// também pode adiar a execução de uma função até o ultimo momento possivel
	defer funcao1()
	funcao2()
	fmt.Println("Student approval Result : ", verifyStudentApproval(7, 8))

}

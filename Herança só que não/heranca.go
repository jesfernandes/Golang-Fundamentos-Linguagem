package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso               string
	instituicaoDeEnsino string
}

func main() {

	pessoa1 := pessoa{"Je", "Fernandes", 23, 170}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Análise e Desenvolvimento de Sistema", "Uni9"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome, estudante1.sobrenome)

}

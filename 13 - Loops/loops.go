package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	//biblioteca time pode ser usada para trabalhar com datas também
	for i < 10 {
		i++
		fmt.Println("Incrementando Variável I = ", i)
		time.Sleep(time.Second)
	}

	for j := 0; j <= 10; j++ {
		fmt.Println("Incrementando Variável J = ", j)
		time.Sleep(time.Second)
	}

	for j := 0; j <= 10; j += 5 {
		fmt.Println("Incrementando Variável J = ", j)
		time.Sleep(time.Second)
	}

	//range
	nomes := [3]string{"Blake", "Douglas", "Margaret"}

	// imprime chave e valor
	for posicao, valor := range nomes {
		fmt.Println(posicao, valor)
	}

	// imprime sem o indice, apenas o valor do array
	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// interando map
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// nao é possivel fazer range no struct (classe)

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

	// Nao existe outro tipo de loop, exemplo o while do Java

}

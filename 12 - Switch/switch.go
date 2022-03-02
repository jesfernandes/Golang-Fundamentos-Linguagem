package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Domingo"
	default:
		return "Valor informado inválido"
	}
}

// O return é passado apenas no final da Função
func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		// falltrough joga o código para dentro da proxima condição(2) sem avaliar a condição
		// nesse caso retorna segunda-feira
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana
}

// Em Go nao existe a cláusula break
// porque ele ja sai automaticamente
// se encontrar uma X condução true
// ele finaliza o switch

func main() {
	dia := diaDaSemana(200)
	fmt.Println(dia)

	fmt.Println("-----------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}

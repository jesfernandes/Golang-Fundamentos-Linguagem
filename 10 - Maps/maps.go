package main

import "fmt"

func main() {

	fmt.Println("Maps")

	//podemos declarar maps chave e valor de tipos diferentes
	//mas todos os itens devem seguir seus tipos correspondentes
	usuario := map[string]string{
		"nome":      "Je",
		"sobrenome": "Fernandes",
	}

	municipios := map[string]map[int]string{
		"zona oeste": {
			1: "Carapicuiba",
			2: "Osasco",
		},
	}

	//para retornar o dado, temos que buscar usando chave
	fmt.Println(usuario["nome"])
	fmt.Println(municipios["zona oeste"])

	//apagando uma chave
	delete(municipios, "zona oeste")
	fmt.Println(municipios)

	//inserindo uma nova chave
	//o array deve seguir os tipos do array principal
	municipios["zona sul"] = map[int]string{
		1: "Jabaquara",
		2: "Guarulhos",
	}

	fmt.Println(municipios)

}

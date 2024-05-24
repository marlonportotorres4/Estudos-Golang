package main

import (
	"fmt"
	
	
	"time"
)

func main() {
	i := 0

	for i < 10 {            //i := 0; i < 10; i++ (for init)
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second) //pausa de 1 segundo
	}

	fmt.Println(i)

	nomes := [3]string{"joao", "lucas", "alfredo"}

	for indice, nome := range nomes { //indice = quantidade de valores, nome = nome do valor
										//colocar "_" no lugar do indice ou nome, para usar apenas um dos dados
	
		fmt.Println(indice, nome)
	}
	for indice, letra := range "PALAVRA"{
	
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome": "Marlon",
		"sobrenome": "Torres",
		
		}

		for chave, valor := range usuario{
		fmt.Println(chave, valor)
		}

		for {
			fmt.Println("Executando infinitamente")
		}

		
}
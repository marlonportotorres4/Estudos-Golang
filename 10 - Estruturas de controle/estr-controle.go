package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Numero maior que 15")
	} else {
		fmt.Println("Numero menor ou igual que 15")
	}

	if outroNumero := numero; outroNumero > 0 { 
		fmt.Println("Maior que zero")
	} else if numero < -10 {
	fmt.Println("menor que -10")
	}else{
	fmt.Println("ENTRE 0 E -10")
	}
 /*Ao criar a variavel "outroNumero" dentro do IF, estamos deixando ela existente
 apenas para aquele escopo do IF, caso tentemos utilizar ela fora do IF, nao conseguiremos.
*/
}

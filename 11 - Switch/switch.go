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
	default:
		return "Número inválido"
	}

}

func diaDaSemana2(numero int) string { //fallthrough - para pular um case
	var diaDaSemana string
	
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda"	
	default:
		diaDaSemana = "Numero invalido"
	}

	return diaDaSemana
}

/*switch { // cada case avalia uma função diferente
case funcao1():
	fmt.Println("funcao1() retornou um valor verdadeiro")
case funcao2():
	fmt.Println("funcao2() retornou um valor verdadeiro")
case funcao3():
	fmt.Println("funcao3() retornou um valor verdadeiro")
} */

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(1)
	fmt.Println(dia)

	dia2 := diaDaSemana2(2)
	fmt.Println(dia2)
}

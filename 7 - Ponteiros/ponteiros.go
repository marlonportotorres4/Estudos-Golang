package main

import "fmt"

func main() {
	fmt.Println("Ponteiro")

	variavel := 10
	variavel2 := variavel //atribuiçao de valor por cópia

	fmt.Println(variavel, variavel2)

	variavel++ //aqui eles ja nao possue mais nenhuma ligação
	fmt.Println(variavel, variavel2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMORIA E NAO UM VALOR
	//BASICAMENTE O ENDEREÇO DE MEMORIA E VÊ O VALOR ALOCADO NELE, NAO O VALOR EM SI
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) //*ponteiro -> para ver o valor do endereço de memoria
									// desreferenciação
	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)

}

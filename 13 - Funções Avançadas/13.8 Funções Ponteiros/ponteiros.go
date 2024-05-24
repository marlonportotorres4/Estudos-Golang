package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int)  {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	//Aqui foi criada uma nova variada e passando um parametro por copia
	//sendo assim o numero  continua 20 e nao -20

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)

	//quando se mexe com ponteiros, todas as alterações vai refletir
	//diretamente na variavel criada
}
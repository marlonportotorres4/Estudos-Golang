package main

import "fmt"

func calculosMatematiocs(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return

	//ACONTECE É QUE O RETORNO JA TA SENDO NOMEADO, ENTAO NAO PRECISAMOS
	//DEFINIR AS VARIAVEIS DENTRO DA FUNÇÃO E NEM COLOCA-LAS NO RETURN
	//APENAS UTILIZALAS NORMALMENTE COMO ESTA ACIMA
}

func main() {
	soma, subtracao := calculosMatematiocs(10, 20)
	fmt.Println(soma, subtracao)
}

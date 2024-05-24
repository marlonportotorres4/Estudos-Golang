package main

import "fmt"

func somar(n1 int, n2 int) int { //PARAMETROS ENTRE() APÓS OS PARENTESES |RETORNO|
	return n1 + n2

}

// mais de um retorno por função
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	var v = func() {
		fmt.Println("Função V")

	}
	v()
	
	var f = func(txt string) string {
		fmt.Println(txt) //AQUI ELE VAI PRINTAR NA TELA IMEDIATAMENTE
		return txt       //AQUI É PRA SER USADO EM OUTRA PARTE
	}

	resultado := f("texto da funcao")
	fmt.Println(resultado)

	resultadoSoma, resultadoSub := calculosMatematicos(50, 40)

	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSub)

	//resultadoSoma, _ -> o "_" serve para ignorar o outro resultado
}

package main

import "fmt"

func funcao1() {
	fmt.Println("Exectando a função 1")
}

func funcao2() {
	fmt.Println("Exectando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("O ALUNO ESTÁ:")
	/*nesse caso, o uuso do defer é para
	a não repetiçao de codigo, portanto essa mensagem ia aparecer
	tanto parra true, quanto para o false, pois so vai serr executada por ultimo*/
	media := (n1 + n2) / 2

	if media >= 7 {
		return true
	}
	return false

}

func main() {
	defer funcao1()
	//DEFER = ADIAR, sempre será executado por ultimo
	funcao2()

	alunoEstaAprovado(10, 7)
}

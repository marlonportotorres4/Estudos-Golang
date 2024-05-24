package main

import "fmt"

func recuperarExec()  {
	if r := recover(); r != nil{
		fmt.Println("Recuperada com suceso")
	}
}


func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExec()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!") //O programa entra em panico e mostra a msg
	//PANIC CHAMA TODAS AS FUNÇÕES QUE TEM DEFER ANTES DE FINALIZAR TUDO
}

func main() {
	fmt.Println(alunoEstaAprovado(6,6))

}
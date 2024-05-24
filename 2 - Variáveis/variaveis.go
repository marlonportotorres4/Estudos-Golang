package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Váriavel 2" //INFERÊNCIA DE TIPO ':='
	var variavel10 = "Oi"
	fmt.Println(variavel10)
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "xxxx"
		variavel4 string = "xxxx"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, varivariavel6 := "zzzzzz", "mmmm"
	fmt.Println(variavel5, varivariavel6)

	//TROCA DE VARIAVEIS:

	variavel1, variavel2 = variavel2, variavel1
	fmt.Println(variavel1, variavel2)

}

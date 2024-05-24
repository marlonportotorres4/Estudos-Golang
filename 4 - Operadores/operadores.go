package main

import "fmt"

func main() {

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	
	//OPERADORES LOGICOS
	fmt.Println(false && true) //AND
	fmt.Println(false || true) //OR
	fmt.Println(!true)         //NO

	//operadores unarios
	numero := 10
	numero++
	numero += 15 // = numero + 15
	numero--
	numero -= 15
	numero *= 3
	numero /= 3
	numero %= 3
}



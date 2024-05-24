package main

import "fmt"

func soma(numeros ...int) int { //receber de 1 a N int.
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int)  {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}



func main() {
	totalDaSoma := soma(1, 2, 3, 4 ,5 , 6, 200, 102, 30, 12)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 10, 9, 2, 3, 4)
}
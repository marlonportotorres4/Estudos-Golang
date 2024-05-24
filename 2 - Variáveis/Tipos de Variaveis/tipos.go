package tipos

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = -10000
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	var numeroReal float64 = 30032323.54
	fmt.Println(numeroReal)

	var palavra string = "oi"
	fmt.Println(palavra)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Mensagem de erro")
	fmt.Println(erro)
}

//int8, int16, int32, int64 -> muda apenas a quantidade de bits que o numero pode ter
//uint -> unsygned int, ou seja, int sem sinal negativo
//rune = int32
//byte = uint8
//float32, float64
//string
//bool - true ou false
//errors.new -> para poder passar uma mensagem no erro

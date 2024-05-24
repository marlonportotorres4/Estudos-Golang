package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa //aqui é onde a herança entra
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	
	pessoa1 := pessoa{"joao", "alexandre", 20, 1.70}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "TI", "uni"}
	fmt.Println(estudante1)
}
			
package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct { //ADICIONANDO STRUCT EM OUTROS STRUCTS --ACIMA--
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var usuario1 usuario
	usuario1.nome = "Marlon"
	usuario1.idade = 19
	usuario1.endereco.logradouro = "rua abc" 
	fmt.Println(usuario1)

	enderecoExemplo := endereco{"alemanha", 30}
	
	usuario2 := usuario{"Kaian", 21, enderecoExemplo}
	fmt.Println(usuario2)

	//quando nao possuo todos os dados
	usuario3 := usuario{nome: "Marlon"}//ou idade: xx (especificando sempre o que voce quer)
	fmt.Println(usuario3)
}
 
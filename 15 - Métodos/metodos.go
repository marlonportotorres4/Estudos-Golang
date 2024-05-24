package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

//METODO é uma ação
func (user usuario) salvar()  {
	fmt.Printf("salvando usuario %s no banco de dados\n", user.nome)
}

func (user usuario) maiorDeIdade() bool  {
	return user.idade>=18
	
}

//metodo para alterar valor dentro do usuario

func (user *usuario) fazerAniversario()  {
	user.idade++
}

func main() {
	fmt.Println("Métodos")
		
	usuario1 := usuario{"joao", 10}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)

}
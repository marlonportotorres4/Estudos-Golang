package main

import (
	"modulo/auxiliar" //importando de outro pacote
	"fmt"
	"github.com/badoux/checkmail" //importando pacotes externos
)
func main() {
	fmt.Println("Olá mundo!")
	
	
	auxiliar.Escrever() //funçao do pacote externo

  	erro := checkmail.ValidateFormat("marlon@gmail.com")
	fmt.Print(erro)


}


/*USAR "GO BUILD" PARA CRIAR E ATUALIZAR O MODULO.exe,
MOSTRANDO AS OUTRAS FUNÇÕES*/

//USAR "GO GET "URL DO SITE" PARA ADICIONAR PACOTE EXTERNO AO GO.MOD "
//"GO MOD TIDY" PARA REMOVER TODAS AS DEPENDENCIAS
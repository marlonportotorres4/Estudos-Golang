package auxiliar

import "fmt"

//Escrever começa com "E" maiusculo para ser usada em outro pacote
func Escrever() {
	fmt.Println("Escrevendo pacote auxiliar")
	escrever2() //funçao do proprio pacote
}
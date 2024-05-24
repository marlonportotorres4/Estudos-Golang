package main

import "fmt"

func main() { //formas diferentes de funções anônimas

	func() {
		fmt.Println("Olá mundo")
	}()

	func (texto string)  {
		fmt.Println(texto)
	}("PASSANDO PARÂMETRO")

	retorno := func (txt string) string  {
		return fmt.Sprintf("Recebido -> %s", txt)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}

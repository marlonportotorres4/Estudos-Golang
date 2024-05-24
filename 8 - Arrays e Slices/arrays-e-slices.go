package main

import (
	"fmt"


)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1" //Acesando as posições
	fmt.Println(array1)

	array2 := [5]int{} //Por Inferência
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	/* SLICE É MUITO MAIS UTILIZADO EM GO AO INVES DO ARRAY
	POIS ELE É MUITO MAIS FLEXIVEL NA DEFINIÇÃO DE VALORES E
	QUANTIDADESS */
	//SLICE NAO É UM ARRAY, APENAS É MUITO SEMELHANTE
	//SAO TIPOS DIFERENTES, SLICE NAO TEM LIMITAÇÕES
	//NAO PODEM INTERAGIR ENTRE SI

	slice1 := []int{20, 30, 40, 50, 60, 70}
	fmt.Println(slice1)

	slice1 = append(slice1, 80) //ADICIONAR VALORES NO SLICE - APPEND
	fmt.Println(slice1)         //Usar o append pela segunrança,, faciidade de uso e flexibilidade
  
	slice2 := array3[1:3] //Pega a fatia do indice 1 e indice 2
	fmt.Println(slice2)

	//ARRAYS INTERNOS

	slice3 := make([]int, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 16)
	slice3 = append(slice3, 18) //SEMPRE QUE A CAPACIDADE É ULTRAPASSADA, AUMENTA A CAPACIDADE
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade
}

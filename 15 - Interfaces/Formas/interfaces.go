package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

/* O QUE ACONTECE?
NÓS QUEREMOS FAZER O CALCULO DA AREA PARA DUAS FORMAS DIFERENTES(POREM ISSO PODERIA VALER PARA MAIS),
ENTAO CRIAMOS DOIS STRUCTS(CIRCULO E QUADRADO) E FIZEMOS UMA INTERFACE CHAMADA "FORMA", ASSIM PODEMOS CHAMAR ESSA INTERFACE POR 
MEIO DE UMA FUNÇÃO "escreverArea" PARA NOS DEVOLVER A AREA DA FORMA QUE QUEREMOS.
E ENTAO CHAMAMOS OS METODOS PARA REALIZARR OS RESPECTIVOS CALCULOS

INTERFACES SAO UTILIZADAS EM CODIGOS COMPLEXOS E DE GRANDE ESCALA*/

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f", f.area())
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func main() {

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}

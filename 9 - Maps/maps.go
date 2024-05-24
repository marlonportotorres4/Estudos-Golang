package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Maps")
	fmt.Println("---------------")

	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)	
	fmt.Println(usuario["nome"])	
		
	usuario2 := map[string]map[string]string {
		
		"nome":{ 
		
			"primeiro": "Joao",
			"segundo": "silva",
		},
		"curso": {
		
			"nome": "Engenharia",
			"campus": "campus 1",
		},
	}

	fmt.Println(usuario2["curso"]["campus"]) //Acessando conteudo do campus no segundo map

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Virgem",
	
	}

fmt.Println(usuario2)

}

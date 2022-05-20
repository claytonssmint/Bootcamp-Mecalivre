package main

import (
	"fmt"
)

var slice = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func quantosAnosTem() {

	qntMaiorQue21Anos := make([]int, 0)

	for key, value := range slice {
		if key == "Benjamin" {
			fmt.Println("A idade é:" )
		}

		if value > 21 {
			qntMaiorQue21Anos = append(qntMaiorQue21Anos)
		}
	}

	fmt.Println("Quantidade de funcionarios com mais de 21 anos é de :", len(qntMaiorQue21Anos))

	slice["Frederico"] = 25
	delete(slice, "Pedro")
}

func main(){
	fmt.Println(((slice)))
}
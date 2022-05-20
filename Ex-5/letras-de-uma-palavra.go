package main

import (
	"fmt"
)

var numero int = 12

func main() {
	for i, letra := range "MERCADOLIVRE" {
		fmt.Println(i, string(letra))
	}

	fmt.Println("Essa palavra contém", numero, "letras")
}

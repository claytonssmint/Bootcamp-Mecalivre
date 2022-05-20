package main

import "fmt"

func main() {
	idade := 30
	empregado := true
	tempoDeAtividadeEmAnos := 2
	salarioAnual := 200.000

	if idade >= 22 && empregado && tempoDeAtividadeEmAnos >= 1 {
		fmt.Println("O cliente está apto a receber um empréstimo")
	} else {
		fmt.Println("O cliente não atende a algum requisito")
	}

	if salarioAnual > 100.000 {
		fmt.Println("O cliente está isento de juros")
	} else {
		fmt.Println("O cliente irá pagar juros sobre o empréstimo")
	}

}
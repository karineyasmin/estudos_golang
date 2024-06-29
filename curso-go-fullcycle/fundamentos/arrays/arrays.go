package main

import "fmt"

var e float64

func main() {
	fmt.Printf("O tipo de E é %T\n", e)

	//arrays em go tem tamanho fixo
	var meuArray [3]int //3 posições fixas
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println((len(meuArray)))
	fmt.Println(meuArray[len(meuArray)-1])

	for indice, valor := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", indice, valor)
	}
}

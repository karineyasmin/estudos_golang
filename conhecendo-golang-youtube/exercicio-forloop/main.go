package main

import (
	"fmt"
	"math/rand"
)

func adivinhaNumero() {
	meuNumero := 35
	numeroTentativas := 0

	for {
		tentativa := rand.Intn(100)

		if tentativa < meuNumero {
			fmt.Println("O número está abaixo", tentativa)
			numeroTentativas++
		} else if tentativa > meuNumero {
			fmt.Println("O número está acima", tentativa)
			numeroTentativas++
		} else {
			fmt.Println("Acertou o número", tentativa)
			break
		}
	}
	fmt.Println("Fim do jogo")
	fmt.Println("Número de tentativas", numeroTentativas)
}

func main() {
	adivinhaNumero()
}

package main //nome do pct

import "fmt" // package com a func Printf

func main() {
	fmt.Printf("Hello, world!\n")

}

// variáveis

var nome string = "Karine" // Posso declarar o valor aqui ou depois
var idade int = 27
var altura float64

// Declaração curta:
// Só deve ser usado dentro de funções || o tipo é inferido automaticamente como em Python =)
// x := 100
// solteira := true

var impressao = fmt.Sprintf("Olá, %s! Você tem %d anos de idade.", nome, idade)

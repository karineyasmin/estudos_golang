// Escreva um programa que solicite ao usuário que insira números inteiros e armazene os números inteiros em uma fatia ordenada.

// O programa deve ser escrito como um loop. Antes de entrar no loop, o programa deve criar uma fatia de número inteiro vazia de tamanho (comprimento) 3.

// Durante cada passagem pelo loop, o programa solicita que o usuário digite um número inteiro a ser adicionado à fatia.

// O programa adiciona o número inteiro à fatia, classifica a fatia e imprime o conteúdo da fatia em ordem de classificação.

// O tamanho do corte deve aumentar para acomodar qualquer número de números inteiros que o usuário decida inserir.

// O programa só deve ser encerrado (saindo do loop) quando o usuário digitar o caractere "X" em vez de um número inteiro.

// Envie seu código-fonte para o programa,
// "slice.go".

package main

import "fmt"

func main() {
	var numbers = [3]string
	var input string

	while(input != "X") {
	fmt.Println("Enter a sequence of whole numbers")
	fmt.Scanln(&input)
	}
}

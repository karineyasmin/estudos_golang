/*

Criar um programa Go de linha de comando onde seja possível fazer o cadastro de notas dos alunos.
O programa deve permitir o cadastro de um aluno com suas 4 notas.
Deve haver uma opção de imprimir uma tabela com os alunos, suas notas e a média.
O programa deverá usar os conceitos vistos até aqui, como: types, slices, for, methods, printf, etc.

	1 - Lançar notas de um aluno
	2 - Imprimir Boletim
	3 - Encerrar
*/

package main

import (
	"fmt"
	"os"
)

type Notas [4]float32
type Boletim map[string]Notas

func adicionarAluno() (string, Notas) {
	var nome string
	var notas Notas

	fmt.Println("Digite o nome do aluno: ")
	fmt.Scanln(&nome)

	for i := 0; i < len(notas); i++ {
		fmt.Printf("Nota %d: ", i+1)
		fmt.Scanln(&notas[i])
	}
	return nome, notas
}

func (notas Notas) calcularMedia() float32 {
	var media float32
	for _, nota := range notas {
		media += nota
	}
	return media / float32(len(notas))
}

func imprimirBoletim(boletim Boletim) {
	fmt.Printf("%-15s % 7s % 7s % 7s % 7s %7s\n", "NOME: ", "AV1", "AV2", "AV3", "AV4", "MÉDIA")
	for nome, notas := range boletim {
		fmt.Printf("%-15s % 7.2f % 7.2f % 7.2f % 7.2f % 7.2f\n",
			nome,
			notas[0],
			notas[1],
			notas[2],
			notas[3],
			notas.calcularMedia(),
		)
	}
}
func main() {

	boletim := Boletim{
		"João":  {7.5, 8.0, 6.5, 7.0},
		"Maria": {6.5, 7.0, 8.0, 7.5},
	}
	for {
		var opcao string // le opcao do menu
		fmt.Println("1 - Adicionar Aluno")
		fmt.Println("2 - Imprimir Boletim")
		fmt.Println("3 - Encerrar")
		fmt.Scanln(&opcao)

		fmt.Print("\n\n")

		switch opcao {
		case "1":
			nome, notas := adicionarAluno()
			boletim[nome] = notas
			fmt.Println("Aluno adicionado com sucesso!")
		case "2":
			imprimirBoletim(boletim)
		case "3":
			os.Exit(0)
		}

	}
}

// Scan -> lê até o primeiro espaço em branco
// Scanln -> lê até a primeira quebra de linha

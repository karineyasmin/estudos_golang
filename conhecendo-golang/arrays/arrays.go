package main

import "fmt"

func zeraNotas(notas [4]float64) {
	for i := range notas {
		notas[i] = 0.0
	}
}

func main() {
	var notas [4]float64
	alunos := [...]string{"João", "Maria"}

	notas[0] = 7
	notas[1] = 8.5
	notas[2] = 9.5
	notas[3] = 10

	fmt.Println(notas[0])
	fmt.Println(notas)
	fmt.Println(alunos)

	media := 0.0
	for _, nota := range notas {
		// fmt.Printf("Verificando a nota %d - %.2f\n", indice+1, nota)
		media += nota
	}

	fmt.Printf("Média: %.2f\n", media/float64(len(notas)))

}

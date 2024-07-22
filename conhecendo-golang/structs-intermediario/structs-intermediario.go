package main

import "fmt"

type aluno struct {
	nome          string
	turma         string
	disciplina    string
	numeroDeNotas uint8
	notas         []float32
}

// NewAluno é um construtor de aluno
func NewAluno(nome string, turma string, disciplina string, notas ...float32) aluno {
	maxNotas := 4
	if len(notas) < maxNotas {
		maxNotas = len(notas)
	}
	return aluno{nome, turma, disciplina, uint8(maxNotas), notas[:maxNotas]}
}

func (a aluno) media() float32 {
	var soma float32

	for _, nota := range a.notas {
		soma += nota
	}
	return soma / float32(a.numeroDeNotas)
}

/*
compara as notas do aluno a com o b
Se a média de a for:
menor: retorna -1
igual: retorna 0
maior: retorna 1
*/
func (a aluno) compare(b aluno) int {
	m1 := a.media()
	m2 := b.media()

	if m1 < m2 {
		return -1
	} else if m1 > m2 {
		return 1
	} else {
		return 0
	}

}
func main() {
	joao := NewAluno("João", "7ª A", "Matemática", 10, 9, 8, 7, 6)
	maria := NewAluno("Maria", "7ª A", "Matemática", 10, 8.3, 7, 5, 6)

	fmt.Println(joao.compare(maria))
}

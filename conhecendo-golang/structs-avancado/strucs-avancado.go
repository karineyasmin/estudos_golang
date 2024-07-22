package main

import "fmt"

type matricula struct {
	id         int
	turma      string
	disciplina string
}

type avaliacoes struct {
	id            int
	numeroDeNotas uint8
	notas         []float32
}

type aluno struct {
	nome       string
	matricula  matricula
	avaliacoes avaliacoes
}

func NewAluno(nome string, turma string, disciplina string, notas ...float32) aluno {
	maxNotas := 4
	if len(notas) < maxNotas {
		maxNotas = len(notas)
	}
	return aluno{nome, matricula{1, turma, disciplina}, avaliacoes{2, uint8(maxNotas), notas[:maxNotas]}}

}

func (a avaliacoes) media() float32 {
	var soma float32

	for _, nota := range a.notas {
		soma += nota
	}
	return soma / float32(a.numeroDeNotas)
}

func main() {
	joao := NewAluno("João", "7ª A", "Matemática", 10, 9, 8, 7, 6)
	maria := NewAluno("Maria", "7ª A", "Matemática", 10, 8.3, 7, 5, 6)

	fmt.Println(maria.matricula.disciplina)
	fmt.Println(joao.matricula.turma)

	fmt.Println(maria.matricula.id)
}

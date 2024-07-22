package main

import (
	"encoding/json"
	"fmt"
)

type aluno struct {
	Nome  string     `json:"nome"`
	Turma string     `json:"turma"`
	Notas [4]float32 `json:"notas"`
}

func main() {
	var boletim []aluno
	var a1 aluno
	var a2 aluno
	var a3 = aluno{Nome: "José", Notas: [4]float32{7.5, 8.0, 9.5, 10.0}}
	// var a4 = a3 // acotnece uma cópia dos valoresN
	a1.Nome = "João"
	a2.Nome = "Maria"
	boletim = append(boletim, a1)
	boletim = append(boletim, a2)
	boletim = append(boletim, a3)

	fmt.Println(boletim)

	a1Json, _ := json.Marshal(a1)

	fmt.Println(string(a1Json))
}

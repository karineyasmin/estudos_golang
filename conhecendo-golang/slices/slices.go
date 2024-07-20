package main

import (
	"fmt"
	"strings"
)

type DiasDaSemana []string

// metodo que modifica o slice original
func (d DiasDaSemana) Uppercase() {
	for i := 0; i < len(d); i++ {
		d[i] = strings.ToUpper(d[i])
	}
}

func main() {
	diasDaSemana := DiasDaSemana{
		"Segunda",
		"Terça",
		"Quarta",
		"Quinta",
		"Sexta",
		"Sábado",
		"Domingo",
	}

	fmt.Println(diasDaSemana)
	diasDaSemana.Uppercase() // mofica o slice original
	fmt.Println(diasDaSemana)

}

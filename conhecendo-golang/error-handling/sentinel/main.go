package main

import (
	"fmt"
)

type MathError string

func (me MathError) Error() string {
	return string(me)

}

const (
	ErrDivisaoPorZero = MathError("Divisão por zero")
	ErrNumeroNegativo = MathError("Não aceita números negativos")
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero
	}
	if a < 0 {
		return 0, ErrNumeroNegativo
	}

	return a / b, nil
}

func main() {
	r, err := divide(10, 0)
	if err != nil {
		if err == ErrDivisaoPorZero {
			fmt.Println("Você não pode dividir por zero")
		} else {
			fmt.Println(r)
		}
	}
}

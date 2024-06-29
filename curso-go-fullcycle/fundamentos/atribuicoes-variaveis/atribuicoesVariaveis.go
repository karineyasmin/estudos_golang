package main // deve sempre ter o nome da pasta ** com exceção do main **

import "fmt"

const a = "Hello, World!"

type ID int // meu próprio tipo

var (
	b bool   = true
	c int    = 10
	d string = "Karine"
	e float64
	f ID      = 1
	g float64 = 12.2
)

// variaveis em escopo local devem ser utilizadas
func main() {
	fmt.Printf("b = %t, c = %d, d = %s, g = %f \n", b, c, d, g)

}

// %T -> exibe o tipo
// %v -> exibe o valor
// %d -> inteiro
// %s -> string
// %f -> float

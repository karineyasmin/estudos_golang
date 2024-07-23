package main

import "fmt"

type Scaler interface {
	Scale(n int)
}

type Coordinate struct {
	x, y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c Coordinate) Scale(n int) {
	c.x *= n
	c.y *= n
}

func main() {

	c1 := Coordinate{2, 4}
	c2 := Coordinate{5, 6}

	coords := []Scaler{c1, c2}

	for _, c := range coords {
		c.Scale(2)
		fmt.Println(c)
	}
}

// & é um ponteiro
// * é um operador de desreferência (pega o valor de um ponteiro)
// %p é um formatador de ponteiro (exibe o endereço de memória)
// maps e arrays sempre são passados como cópia

package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 4, 5, 76, 43, 12, 65, 87, 32, 1, 65, 32))
	fmt.Println(sum(1, 4, 5))

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// ... -> n parametros

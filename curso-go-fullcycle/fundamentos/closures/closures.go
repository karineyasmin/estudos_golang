package main

import "fmt"

func main() {

	total := func() int {
		return sum(1, 4, 5, 76, 43, 12, 65, 87, 32, 1, 65, 32) * 2
	}()

	fmt.Println(total)

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// funcao anonima
// fun() {

// }()

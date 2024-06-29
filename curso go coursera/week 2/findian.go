/*
Escreva um programa que solicite ao usuário que digite uma string.

O programa procura os caracteres 'i', 'a' e 'n' na string inserida.

O programa deve imprimir "Found!" se a cadeia de caracteres inserida começar com o caractere "i", terminar com o caractere "n" e contiver o caractere "a".
Caso contrário, o programa deve imprimir "Not Found!". O programa não deve diferenciar maiúsculas de minúsculas, portanto, não importa se os caracteres são maiúsculos ou minúsculos.

Exemplos: O programa deve imprimir "Found!" para as seguintes cadeias de caracteres inseridas, "ian", "Ian", "iuiygaygn", "I d skd a efju N".
O programa deve imprimir "Not Found!" para as seguintes cadeias de caracteres, "ihhhhhn", "ina", "xian".

Envie seu código-fonte para o programa,
"findian.go".


*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Insert a String: ")

	fmt.Scanln(&input)

	input = strings.ToLower(input)
	fmt.Println(input)
	foundI := strings.HasPrefix(input, "i")
	foundA := strings.Contains(input, "a")
	foundN := strings.HasSuffix(input, "n")

	if foundI && foundA && foundN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

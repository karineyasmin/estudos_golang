package main

import "fmt"

// * maps -> chave e valor
// map[tipo da chave]tipo do valor{}
func main() {

	salarios := map[string]int{"João": 1000, "Alice": 2000, "Alberto": 1500, "Maria": 4000}
	fmt.Println(salarios["Alice"])
	delete(salarios, "Alberto")
	salarios["Júlia"] = 1350
	fmt.Println(salarios["Júlia"])
	fmt.Println(salarios)

	// sal := make(type, 0)
	// new_map := map[string]int{} // inicia map vazio

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	// _ -> blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salário é de %d\n", salario)

	}

}

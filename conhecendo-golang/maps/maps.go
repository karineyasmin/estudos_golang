package main

import "fmt"

func dobraPopulacao(p map[string]int) {
	for key, value := range p {
		p[key] = value * 2
	}
}

func main() {
	var idadeAlunos map[string]int
	// Inicializando o map
	idadeAlunos = map[string]int{
		"João": 25,
	}
	idadeAlunos["André"] = 27
	idadeAlunos["Caetano"] = 30
	idadeAlunos["Caetano"] = 31

	fmt.Println(idadeAlunos)

	andre, ok := idadeAlunos["André"]
	fmt.Printf("A idade de André: %d\n", andre)
	println(ok)

	delete(idadeAlunos, "André")
	fmt.Println(idadeAlunos)

	populacaoPorUf := make(map[string]int, 27)
	populacaoPorUf["SP"] = 12_000_000
	populacaoPorUf["RJ"] = 6_000_000
	fmt.Println(populacaoPorUf)
	fmt.Println(len(populacaoPorUf))

	for uf, populacao := range populacaoPorUf {
		fmt.Println(populacaoPorUf[uf])
		fmt.Printf("A população de %s é %d\n", uf, populacao)
	}

	dobraPopulacao(populacaoPorUf)
	fmt.Println(populacaoPorUf)
}

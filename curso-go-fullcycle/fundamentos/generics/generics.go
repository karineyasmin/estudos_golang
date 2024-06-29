package main

// Generics são um recurso que permite a criação de funções, estruturas e interfaces que aceitam qualquer tipo de dado.

// Exemplo de função genérica
func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}
func main() {
	m := map[string]int{"Karine": 1000, "Yasmin": 2000, "Fernanda": 3000}
	println(SomaInteiro(m))
}

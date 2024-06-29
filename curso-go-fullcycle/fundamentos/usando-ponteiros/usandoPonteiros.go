package main

// se usar o * ele altera os valores e não uma cópia dele

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	soma(&minhaVar1, &minhaVar2) //referencia da memoria
	println(minhaVar1)           // imprime 10 porque a variável minhaVar1 não foi alterada

}

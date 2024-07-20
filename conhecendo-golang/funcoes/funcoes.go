package main

import "fmt"

func hello(name string) {
	fmt.Println("Hello,", name)
}

func soma(a int, b int) (int, string) {
	result := a + b
	resultFmt := fmt.Sprintf("%d + %d = %d", a, b, result) // Sprintf returns a string
	return result, resultFmt
}

func subtracao(a int, b int) (c int, d string) {
	defer fmt.Println("Fim da execução da subtração") // defer is executed after the function returns
	fmt.Println("Início da execução da subtração")
	c = a + b
	d = fmt.Sprintf("%d - %d = %d", a, b, c)
	return // c and d are returned
}

func counter() {
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
}

func main() {
	resultado, resultadoFormatado := soma(3, 5)
	fmt.Println(resultado)
	fmt.Println(resultadoFormatado)

	resultadoSub, resultadoSubFmt := subtracao(10, 5)
	fmt.Println(resultadoSub)
	fmt.Println(resultadoSubFmt)

	counter()
}

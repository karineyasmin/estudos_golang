package main

import (
	"fmt"
)

type funcao func(string) // type alias para funcao

func hello(name string) {
	fmt.Println("Hello", name)
}

func funcaoComoParametro(f funcao) { // funcao como parametro usando type alias
	f("Karine")

}

func funcaoQueRetornaFuncao(s string) func(string) {
	return func(name string) {
		fmt.Println(s, name)
	}

}

func somador() func(int) int {
	soma := 0
	return func(valor int) int {
		soma += valor
		return soma

	}
}

func main() {
	var h = hello //atribui a funcao hello a variavel h
	h("Karine")   //chama a funcao hello atraves da variavel h

	func() {
		fmt.Println("Função anônima")
	}() //cria e chama uma função anônima

	funcaoComoParametro(hello) //passa o nome da funcao hello como parametro

	s := funcaoQueRetornaFuncao("Hello")
	s("Karine")

	soma := somador()
	soma(5)
	soma(10)
	fmt.Println(soma(10))
}

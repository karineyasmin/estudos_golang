package main

import "fmt"

func main() {
	fmt.Println("Interfaces vazias são interfaces que não possuem métodos.")

	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}

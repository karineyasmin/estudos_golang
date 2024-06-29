package main

import "fmt"

func main() {

	var minhaVar interface{} = "Karine Yasmin"

	println(minhaVar.(string))

	resultado, ok := minhaVar.(int) // caso a coversao não seja possível, resultado será 1 e ok será false
	println(resultado, ok)

}

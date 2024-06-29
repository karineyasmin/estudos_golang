package main

func main() {

	// Memoria -> endereço -> valor
	// & -> endereço
	//* -> ponteiro

	a := 10
	println(&a)            // endereço de a
	var ponteiro *int = &a // ponteiro para o endereço de a
	println(ponteiro)      // endereço de a

	*ponteiro = 20 // valor de a
	b := &a        // endereço de a
	println(*b)    // valor de a
	*b = 30
	println(a)

}

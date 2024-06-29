// Escreva um programa que solicite ao usuário que primeiro digite um nome e depois um endereço.

// Seu programa deve criar um mapa e adicionar o nome e o endereço ao mapa usando as chaves "name" e "address", respectivamente.

// Seu programa deve usar Marshal() para criar um objeto JSON a partir do mapa e, em seguida, deve imprimir o objeto JSON.

// Envie seu código-fonte para o programa,
// "makejson.go".

//-------------------------------------------------------------------------------------------------------------------------------------

package main

import "fmt"

func main() {

	var (
		inputName   string
		inputAdress string
	)

	println("Enter a name:")
	fmt.Scan(&inputName)
	println("Enter a adress:")
	fmt.Scan(&inputAdress)

	mapJson := map[string]string{"name": inputName, "adress": inputAdress}

	mapJson["adress"] = inputAdress
	mapJson["name"] = inputName

	fmt.Println(mapJson)
}


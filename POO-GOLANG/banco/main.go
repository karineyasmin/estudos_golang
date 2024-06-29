package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaKarine := contas.ContaCorrente{titular: "Karine", saldo: 300}
	contaFernanda := contas.ContaCorrente{titular: "Fernanda", saldo: 100}

	status := contaKarine.Transferir(200, &contaFernanda)
	fmt.Println(status)
	fmt.Println(contaKarine)
	fmt.Println(contaFernanda)
}

// outras maneiras de implentar a struct

// contaFernanda := ContaCorrente{"Fernanda", 222, 234567, 100.5}
// fmt.Println(contaFernanda)

// var contaHenry *ContaCorrente
// contaHenry = new(ContaCorrente)
// contaHenry.titular = "Henry"
// contaHenry.numeroAgencia = 111
// contaHenry.numeroConta = 32541
// contaHenry.saldo = 200.0

// fmt.Println(*contaHenry)

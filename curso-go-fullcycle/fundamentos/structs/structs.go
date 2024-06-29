package main

import "fmt"

type Pessoa interface {
	// interface do go permite somente MÉTODOS
	Desativar()
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // composição
}

// Método
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("A cliente %s foi desativada", c.Nome)

}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}
func main() {

	karine := Cliente{
		Nome:  "Karine",
		Idade: 27,
		Ativo: true,
	}

	karine.Ativo = false
	// karine.Estado = "SP"
	karine.Desativar()
	fmt.Println(karine)
	Desativacao(karine)
}

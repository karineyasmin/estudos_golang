package main

import (
	"errors"
	"fmt"
	"os"
)

func verticaArquivo(nome string) error {
	f, err := os.Open(nome)
	if err != nil {
		return fmt.Errorf("Erro ao abrir o arquivo: %w", err)
	}
	f.Close()
	fmt.Println("Arquivo existe")
	return nil
}

func main() {
	err := verticaArquivo("arquivo.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println(errors.Unwrap((err)))
	}

}

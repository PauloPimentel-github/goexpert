package main

import (
	"fmt"

	"curso-goexpert/fundacao/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Honda"}
	fmt.Println("Resultado: ", soma)
	fmt.Println(matematica.A)
	fmt.Println(carro)
	fmt.Println(carro.Andar())
}

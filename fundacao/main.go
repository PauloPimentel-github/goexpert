package main

import (
	"curso-goexpert/fundacao/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Honda"}
	fmt.Println("Resultado: ", soma)
	fmt.Println(matematica.A)
	fmt.Println(carro)
	fmt.Println(carro.Andar())
	fmt.Println(uuid.New())
}

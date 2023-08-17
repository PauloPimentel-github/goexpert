package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	ph := Cliente{
		Nome:  "PH",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", ph.Nome, ph.Idade, ph.Ativo)

	ph.Ativo = false
	fmt.Println(ph.Ativo)
}

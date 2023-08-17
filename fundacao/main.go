package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	ph := Cliente{
		Nome:  "PH",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", ph.Nome, ph.Idade, ph.Ativo)

	ph.Ativo = false
	ph.Endereco.Cidade = "São Paulo"
	fmt.Println(ph.Ativo)
}

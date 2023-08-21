package main

import "fmt"

type Pessoa interface {
	Desativar()
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Empresa struct {
	Nome string
}

func (empresa Empresa) Desativar() {

}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	ph := Cliente{
		Nome:  "PH",
		Idade: 30,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}

	Desativacao(ph)
	Desativacao(minhaEmpresa)
}

package main

import "fmt"

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (conta *Conta) simular(valor int) int {
	conta.saldo += valor
	println(conta.saldo)
	return conta.saldo
}

func main() {
	conta := Conta{
		saldo: 100,
	}
	conta.simular(200)
	fmt.Printf("O valor da struct com saldo %v", conta.saldo)
}

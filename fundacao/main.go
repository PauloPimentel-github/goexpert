package main

func main() {

	// Memória -> Endereço -> Valor
	// Variável -> ponteiro que tem um endereço na memória -> valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	println(ponteiro)
	println(a)
	println(b)
	println(*b)

	*b = 30
	println(a)
}

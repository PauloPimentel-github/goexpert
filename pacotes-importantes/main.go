package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//criando arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//escrevendo no arquivo
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	// tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	//leitura no arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	//remover arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}

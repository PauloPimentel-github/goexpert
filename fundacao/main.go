package main

import "fmt"

func main() {
	salarios := map[string]int{"PH": 1000, "Pimentel": 2000, "Henrique": 3000}
	fmt.Println(salarios["PH"])

	delete(salarios, "PH")
	fmt.Println(salarios)

	salarios["PH"] = 5000
	fmt.Println(salarios["PH"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Luiz"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}

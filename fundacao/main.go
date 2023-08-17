package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, nummero := range numeros {
		total += nummero
	}
	return total
}

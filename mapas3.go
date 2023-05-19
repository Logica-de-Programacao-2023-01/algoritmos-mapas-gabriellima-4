// Escreva uma função que receba um mapa com valores inteiros e retorne a soma de todos os valores.
package main

import "fmt"

func somarEmMapa(mapa map[string]int) int {
	soma := 0

	for _, valor := range mapa {
		soma += valor
	}
	return soma
}

func main() {
	mapa := map[string]int{
		"n1": 12,
		"n2": 30,
		"n3": 284,
	}
	fmt.Println(somarEmMapa(mapa))
}

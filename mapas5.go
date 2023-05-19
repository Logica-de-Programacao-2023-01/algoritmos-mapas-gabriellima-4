// Escreva uma função que receba uma string e retorne um mapa onde as chaves são os caracteres presentes na string
// e os valores são a frequência de cada caractere.
package main

import (
	"fmt"
	"strings"
)

func contarLetras(s string) map[string]int {
	mapa := make(map[string]int)
	caracteres := strings.Split(s, "")

	for _, letras := range caracteres {
		mapa[letras] += 1
	}

	return mapa
}

func main() {
	str := "To precisando de uma praia."

	fmt.Println(contarLetras(str))
}

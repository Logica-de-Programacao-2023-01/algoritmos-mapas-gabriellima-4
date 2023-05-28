// Escreva uma função que receba uma lista de mapas, onde cada mapa contém a contagem de palavras de um texto,
// e retorne um único mapa contendo a soma de todas as contagens.
package main

import (
	"fmt"
	"strings"
)

func contarPalavras2(s string) map[string]int {
	palavras := strings.Fields(s)
	mapa := make(map[string]int)

	for _, palavra := range palavras {
		mapa[palavra]++
	}

	return mapa
}

func juntarmapas2(mapas []map[string]int) map[string]int {
	mapa1 := make(map[string]int)

	for _, rangemapa := range mapas {
		for chave, valor := range rangemapa {
			mapa1[chave] += valor
		}
	}
	return mapa1
}

func main() {
	s1 := "Do jeito que a vida quer"
	s2 := "Pé na Areia"
	s3 := "Deixa a vida me levar"

	mix := []map[string]int{contarPalavras2(s1), contarPalavras2(s2), contarPalavras2(s3)}
	fmt.Println(juntarmapas2(mix))
}

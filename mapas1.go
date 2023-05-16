// Escreva uma função que conte a ocorrência de cada palavra em uma string e
// retorne um mapa onde as chaves são as palavras encontradas
// e os valores são o número de ocorrências de cada palavra.
package main

import (
	"fmt"
	"strings"
)

func contarPalavras(texto string) map[string]int {
	palavras := strings.Fields(texto)
	ocorrencias := make(map[string]int)
	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}
	return ocorrencias
}

func main() {
	//Últimos 10 campeões brasileiros
	x := "Palmeiras , Atlético , Flamengo , Flamengo , Palmeiras , Corinthians , Palmeiras , Corinthians ," +
		" Cruzeiro , Cruzeiro"
	fmt.Println(contarPalavras(x))
}

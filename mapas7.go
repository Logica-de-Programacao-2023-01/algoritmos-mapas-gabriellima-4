//Escreva uma função que receba uma string contendo uma frase e retorne um mapa onde as chaves
//são as palavras encontradas na frase e
//os valores são mapas contendo a contagem de cada letra em cada palavra.
package main

import (
	"fmt"
	"strings"
)

func contarLetrasEPalavras(s string) map[string]map[string]int {
	mapaResultante := make(map[string]map[string]int)
	palavras := strings.Fields(s)

	caracteres := []string{}

	for _, palavra := range palavras {
		caracteres = strings.Split(palavra, "")

		mapaResultante[palavra] = make(map[string]int)

		for _, letra := range caracteres {
			mapaResultante[palavra][letra]++
		}

	}

	return mapaResultante
}

func main() {
	s := "Ayrton Senna do Brasil"
	fmt.Println(contarLetrasEPalavras(s))
}

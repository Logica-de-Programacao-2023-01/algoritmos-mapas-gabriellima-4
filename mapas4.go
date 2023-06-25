//Escreva uma função que receba um slice de palavras e retorne um mapa onde as chaves são palavras que são anagramas
//entre si e os valores são os grupos de palavras anagramas.
package main

import (
	"fmt"
	"sort"
	"strings"
)

func acharAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {

		sorted := sortString(palavra)

		anagramas[sorted] = append(anagramas[sorted], palavra)
	}

	return anagramas
}

func sortString(s string) string {

	letras := strings.Split(s, "")

	sort.Strings(letras)

	return strings.Join(letras, "")
}

func main() {

	nomes := []string{"gabriel","julia","arthur","bruno","anna","lucas","alice","martin"}

	fmt.Println(acharAnagramas(nomes))

}

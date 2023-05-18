// Escreva uma função que receba dois mapas e retorne um novo mapa contendo todos os elementos dos mapas de entrada.
// Em caso de chaves duplicadas, o valor do segundo mapa deve prevalecer.
package main

import "fmt"

func juntarMapas(m1 map[string]int, m2 map[string]int) map[string]int {

	mapaResultante := make(map[string]int)

	for chave, valor := range m1 {
		mapaResultante[chave] = valor
	}
	for chave, valor := range m2 {
		mapaResultante[chave] = valor
	}
	
	return mapaResultante
}

func main() {
	// seeds para os NBA Playoffs
	mapa1 := map[string]int{
		"DEN": 1,
		"PHX": 4,
		"GSW": 6,
		"LAL": 7,
	}

	mapa2 := map[string]int{
		"MIA": 8,
		"NYK": 5,
		"PHI": 2,
		"BOS": 3,
	}

	fmt.Println(juntarMapas(mapa1, mapa2))
}

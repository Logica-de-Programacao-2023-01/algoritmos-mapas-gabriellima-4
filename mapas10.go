// Escreva uma função que receba um slice de inteiros e retorne um mapa
// onde as chaves são os pares de números encontrados no slice e os valores são as frequências de cada par.
package main

import "fmt"

func Pares(slice []int) map[int]int {
	frequenciaS := make(map[int]int)

	for _, inteiro := range slice {
		frequenciaS[inteiro] += 1
	}
	for numero, frequencia := range frequenciaS {

		if frequencia%2 != 0 {
			delete(frequenciaS, numero)
		} else {
			frequenciaS[numero] /= 2
		}
	}
	return frequenciaS
}

func main() {
	slice := []int{1, 4, 3, 6, 7, 2, 3, 1, 8, 9, 0, 0, 2, 7, 3}

	fmt.Println(Pares(slice))
}

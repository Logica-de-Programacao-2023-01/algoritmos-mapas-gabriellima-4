// Escreva uma função que receba um mapa onde as chaves são nomes de pessoas
// e os valores são as quantias de dinheiro que cada pessoa gastou em uma conta compartilhada.
// A função deve calcular o valor que cada pessoa deve receber ou pagar para igualar as despesas.
package main

import "fmt"

func acertoDeContas(mapa map[string]float64) {
	maiorQuantia := 0.0

	for _, valor := range mapa {
		if valor > maiorQuantia {
			maiorQuantia = valor
		}
	}
	for nome, conta := range mapa {
		conta = maiorQuantia - conta
		fmt.Printf("A pessoa %s deve pagar %f para acertar as contas.\n", nome, conta)
	}

}

func main() {
	mapa := map[string]float64{
		"Gabriel": 1000,
		"Laura":   380,
		"Anna":    250,
	}
	acertoDeContas(mapa)
}

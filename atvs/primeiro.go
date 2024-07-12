package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "   OLA mundo!    "
	primeiroCaracter := frase[0]//imprimir uma carcter
	fmt.Printf("Primeira caracter: %c", primeiroCaracter)
	segundaFrase := frase[4:9]//imprimir uma quantidade de caracter
	fmt.Printf("\nSegunda frase: %s", segundaFrase)
	frasesminusculo := strings.ToLower(frase)//Converter para minusculo
	fmt.Printf("\nFrase minusculo: %s", frasesminusculo)
	fraseMaiusculo := strings.ToUpper(frase)//Converter para maiusculo
	fmt.Printf("\nFrase Maiusculo: %s", fraseMaiusculo)
	tirarEspaços := strings.Trim(frase, " ")
	fmt.Printf("\nTirando Espaços: %s", tirarEspaços)
	quantasCaracter := len(frase)//ler quanto de espaço tem ultilozado
	fmt.Printf("\nQuantos Espaços tem: %d", quantasCaracter)

}

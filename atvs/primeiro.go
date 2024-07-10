package main

import "fmt"

func main() {
	var nome string
	fmt.Print("Qual o nome do usuario: ")
	fmt.Scanf("%s", &nome)
	fmt.Printf("O nome do usuario é %s!\n", nome) 

	texto := "Qual o nome da serie: " 
	fmt.Print(texto)
	var serie string
	fmt.Scanf("%s", &serie)
	fmt.Printf("O nome da serie é %s\n", serie) 
}

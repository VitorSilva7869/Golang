package main

import "fmt"

func main() {
    //condiçoes 
	idade := 25

    if idade < 13 {
        fmt.Println("Você é uma criança.")
    } else if idade < 20 {
        fmt.Println("Você é um adolescente.")
    } else if idade < 60 {
        fmt.Println("Você é um adulto.")
    } else {
        fmt.Println("Você é um idoso.")
    }
}

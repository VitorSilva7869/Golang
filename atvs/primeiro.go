package main

import "fmt"

func main() {
	//Loop for básico	
	for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
	//Loop for estilo while
	i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
	//Loop for infinito
    for {
        fmt.Println(i)
        i++
        if i >= 5 {
            break // Usamos 'break' para sair do loop
        }
    }
	//Iterar sobre um array ou slice
	numeros := []int{1, 2, 3, 4, 5}

    for index, value := range numeros {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
	//Iterar sobre um mapa
	capitals := map[string]string{
        "France": "Paris",
        "Italy":  "Rome",
        "Japan":  "Tokyo",
    }

    for country, capital := range capitals {
        fmt.Printf("The capital of %s is %s\n", country, capital)
    }
	//Iterar sobre uma string
	message := "Hello, World!"

    for index, char := range message {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }
	//Uso de continue e break
	for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Pula números pares
        }
        if i > 7 {
            break // Sai do loop quando i for maior que 7
        }
        fmt.Println(i)
    }
}

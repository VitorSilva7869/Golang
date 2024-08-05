package main

import "fmt"

func main() {
    //-----Array-------
    //Definindo e Inicializando um Array
    var arr1 [5]int // Array de inteiros com 5 elementos
    arr2 := [3]string{"Go", "Python", "Java"} // Array de strings com 3 elementos

    fmt.Println(arr1) // Output: [0 0 0 0 0]
    fmt.Println(arr2) // Output: [Go Python Java]

    //Acessando e Modificando Elementos de um Array
    arr := [3]int{10, 20, 30}
    fmt.Println(arr[0]) // Output: 10

    arr[1] = 40
    fmt.Println(arr) // Output: [10 40 30]

    //--------Slice---------

    //Definindo e Inicializando um Slice
    var slice1 []int // Slice de inteiros
    slice2 := []string{"Go", "Python", "Java"} // Slice de strings

    fmt.Println(slice1) // Output: []
    fmt.Println(slice2) // Output: [Go Python Java]

    //Criando um Slice a partir de um Array
    arr4 := [5]int{10, 20, 30, 40, 50}
    slice := arr4[1:4] // Cria um slice do índice 1 ao 3 (não inclui o 4)

    fmt.Println(slice) // Output: [20 30 40]

    //Modificando Elementos de um Slice
    slice3 := []int{10, 20, 30}
    slice3[1] = 40
    fmt.Println(slice3) // Output: [10 40 30]

    //append-- Adiciona elementos a um slice. Se necessário, aloca um novo array para acomodar os novos elementos.
    slice5 := []int{10, 20, 30}
    slice5 = append(slice5, 40, 50)
    fmt.Println(slice5) // Output: [10 20 30 40 50]
    
    //len -- Retorna o comprimento (número de elementos) do slice.
    slice4 := []int{10, 20, 30}
    fmt.Println(len(slice4)) // Output: 3

    //cap -- Retorna a capacidade do slice, ou seja, o número de elementos que o slice pode armazenar sem precisar de uma nova alocação.
    slice6 := make([]int, 3, 5) // Cria um slice com comprimento 3 e capacidade 5
    fmt.Println(cap(slice6))    // Output: 5

}

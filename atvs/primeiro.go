package main

import (
	"fmt"
	"math"
)

func main() {
	// uint32 -> 0 a 4.294.967.295
	// uint64 -> 0 a 18.446.744.073.709.551.615
	// int32 -> -2.147.483.648 a 2.147.483.647
	// int64 -> -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807
	var quantidade uint = 24
	var termperatura int = -3
	fmt.Printf("Quantidade : %d e Temperatura: %d", quantidade, termperatura)
	var numero float64 = 1.5
	fmt.Printf("\nQuantidade: %.0f", numero)//%.0f Aproxima. %0.2f vai ate duas casas decimais

	//imc
	fmt.Println("Digite a sua altura:")
    var altura float64
    fmt.Scanf("%f\n", &altura)

    fmt.Println("Digite o seu peso:")
    var peso float64
    fmt.Scanf("%f\n", &peso)

    imc := peso / math.Pow(altura, 2)//potencia 
    fmt.Printf("O seu IMC: %f\n", imc)

	// Operadores básicos
    a := 8
    b := 3

    sum := a + b
    difference := a - b
    product := a * b
    quotient := a / b
    remainder := a % b

    fmt.Println("Soma:", sum)
    fmt.Println("Diferença:", difference)
    fmt.Println("Produto:", product)
    fmt.Println("Quociente:", quotient)
    fmt.Println("Resto:", remainder)

    // Funções da biblioteca math
    sqrt := math.Sqrt(16)
    power := math.Pow(2, 3)
    sin := math.Sin(math.Pi / 2)
    cos := math.Cos(0)
    tan := math.Tan(math.Pi / 4)
    log := math.Log(1)
    abs := math.Abs(-5)

    fmt.Println("Raiz quadrada de 16:", sqrt)
    fmt.Println("2 elevado a 3:", power)
    fmt.Println("Seno de π/2:", sin)
    fmt.Println("Cosseno de 0:", cos)
    fmt.Println("Tangente de π/4:", tan)
    fmt.Println("Logaritmo natural de 1:", log)
    fmt.Println("Valor absoluto de -5:", abs)

}

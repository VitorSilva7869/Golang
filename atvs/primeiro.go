package main

import "fmt"

func main() {
    var valor1 bool
    var valor2 bool
    valor1 = true
    valor2 = true

    resultado := valor1 && valor2
    fmt.Printf("%t && %t = %t\n", valor1, valor2, resultado)

    resultadoAnd1 := true && true
    resultadoAnd2 := true && false
    resultadoAnd3 := false && true
    resultadoAnd4 := false && false

    fmt.Printf("true && true = %t\n", resultadoAnd1)
    fmt.Printf("true && false = %t\n", resultadoAnd2)
    fmt.Printf("false && true = %t\n", resultadoAnd3)
    fmt.Printf("false && false = %t\n", resultadoAnd4)

	resultado01 := valor1 || valor2
    fmt.Printf("%t || %t = %t\n", valor1, valor2, resultado01)

    resultadoOr1 := true || true
    resultadoOr2 := true || false
    resultadoOr3 := false || true
    resultadoOr4 := false || false

    fmt.Printf("\ntrue || true = %t\n", resultadoOr1)
    fmt.Printf("true || false = %t\n", resultadoOr2)
    fmt.Printf("false || true = %t\n", resultadoOr3)
    fmt.Printf("false || false = %t\n", resultadoOr4)

    negacaoValor1 := !valor1
    fmt.Printf("\n\nNegação de %t: %t\n", valor1, negacaoValor1)

	numero1 := 7
    numero2 := 10

    resultado1 := numero1 == numero2
    resultado2 := numero1 > numero2
    resultado3 := numero1 < numero2
    resultado4 := numero1 >= numero2
    resultado5 := numero1 <= numero2
    resultado6 := numero1 != numero2

    fmt.Printf("%d == %d = %t\n", numero1, numero2, resultado1)
    fmt.Printf("%d > %d = %t\n", numero1, numero2, resultado2)
    fmt.Printf("%d < %d = %t\n", numero1, numero2, resultado3)
    fmt.Printf("%d >= %d = %t\n", numero1, numero2, resultado4)
    fmt.Printf("%d <= %d = %t\n", numero1, numero2, resultado5)
    fmt.Printf("%d != %d = %t\n", numero1, numero2, resultado6)
}

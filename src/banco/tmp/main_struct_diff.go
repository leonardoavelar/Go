package main

import (
	"fmt"	
)

type ContaCorrente struct {
	titular string
	agencia int
	conta int
	saldo float64
}

func main() {

	leonardo := ContaCorrente { "Leonardo Avelar", 1234, 5678, 57020.50 }
	leonardo2 := ContaCorrente { "Leonardo Avelar", 1234, 5678, 57020.50 }

	imprimir(leonardo)
	println(leonardo == leonardo2)
	println(&leonardo)
	println(&leonardo2)

	fernanda := ContaCorrente { titular: "Fernanda", saldo: 23900.89}
	fernanda2 := ContaCorrente { titular: "Fernanda", saldo: 23900.89}

	imprimir(fernanda)
	println(fernanda == fernanda2)

	matheus := *new(ContaCorrente)
	matheus.titular = "Matheus"
	matheus.saldo = 1290.58

	matheus2 := *new(ContaCorrente)
	matheus2.titular = "Matheus"
	matheus2.saldo = 1290.58

	imprimir(matheus)
	println(matheus == matheus2)

	var daniel ContaCorrente
	daniel = *new(ContaCorrente)
	daniel.titular = "Daniel"
	daniel.agencia = 1455
	daniel.conta = 4353
	daniel.saldo = 3455.01

	var daniel2 ContaCorrente
	daniel2 = *new(ContaCorrente)
	daniel2.titular = "Daniel"
	daniel2.agencia = 1455
	daniel2.conta = 4353
	daniel2.saldo = 3455.01

	imprimir(daniel)
	println(daniel == daniel2)

}

func imprimir(value ContaCorrente) {

	println("Titular:", value.titular, "| Agencia:", value.agencia, "| Conta:", value.conta, "| Saldo:", fmt.Sprintf("%.2f", value.saldo))

}
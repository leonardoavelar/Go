package main

import (
	"fmt"	
)

type contaCorrente struct {
	titular string
	agencia int
	conta int
	saldo float64
}

func main() {

	leonardo := contaCorrente { "Leonardo Avelar", 1234, 5678, 57020.50 }

	fernanda := contaCorrente { titular: "Fernanda", saldo: 23900.89}

	matheus := *new(contaCorrente)
	matheus.titular = "Matheus"
	matheus.saldo = 1290.58

	var daniel contaCorrente
	daniel = *new(contaCorrente)
	daniel.titular = "Daniel"
	daniel.agencia = 1455
	daniel.conta = 4353
	daniel.saldo = 3455.01

	imprimir(leonardo)
	imprimir(fernanda)
	imprimir(matheus)
	imprimir(daniel)

	pix(&leonardo, &fernanda, 1000)
	leonardo.pix(&fernanda, 2000)
	matheus.pix(&daniel, 3000)
	leonardo.pix(&matheus, 2000)
	daniel.pix(&leonardo, -1000)

	imprimir(leonardo)
	imprimir(fernanda)
	imprimir(matheus)
	imprimir(daniel)

}

func imprimir(value contaCorrente) {

	// println("---------------------------------------------")
	println("Titular:", value.titular, "| Agencia:", value.agencia, "| Conta:", value.conta, "| Saldo:", fmt.Sprintf("%.2f", value.saldo))

}

func pix(saque *contaCorrente, deposito *contaCorrente, valor float64) {

	println("---------------------------------------------")

	if valor < 0 {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do pix é negativo")

	} else if saque.saldo < valor {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do pix é superior ao saldo da conta do cliente:", saque.titular)	
	
	} else {

		saque.saldo -= valor
		deposito.saldo += valor

		println("Transação realizado com sucesso")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("Saque realizado na conta do cliente:", saque.titular)
		println("Depósito realizado na conta do cliente:", deposito.titular)	

	}

	println("---------------------------------------------")

}

func (t *contaCorrente) pix(deposito *contaCorrente, valor float64) {

	println("---------------------------------------------")

	if valor < 0 {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do pix é negativo")

	} else if t.saldo < valor {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do pix é superior ao saldo da conta do cliente:", t.titular)	
	
	} else {

		t.saldo -= valor
		deposito.saldo += valor

		println("Transação realizado com sucesso")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("Saque realizado na conta do cliente:", t.titular)
		println("Depósito realizado na conta do cliente:", deposito.titular)	

	}

	println("---------------------------------------------")

}
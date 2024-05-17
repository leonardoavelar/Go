package contas

import (
	"fmt"
	"banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Cliente
	Agencia int
	Conta int
	saldo float64
}

func (t *ContaCorrente) Saldo() float64 {
	return t.saldo
}

func (t *ContaCorrente) Imprimir() {

	// println("---------------------------------------------")
	println(t.Titular.Descricao(), "| Agencia:", t.Agencia, "| Conta:", t.Conta, "| saldo:", fmt.Sprintf("%.2f", t.Saldo()))

}

// func Pix(saque *ContaCorrente, deposito *ContaCorrente, valor float64) {

// 	println("---------------------------------------------")

// 	if valor < 0 {

// 		println("Transação cancelada")
// 		println("Valor:", fmt.Sprintf("%.2f", valor))
// 		println("O valor do pix é negativo")

// 	} else if saque.saldo < valor {

// 		println("Transação cancelada")
// 		println("Valor:", fmt.Sprintf("%.2f", valor))
// 		println("O valor do pix é superior ao saldo da Conta do cliente:", saque.Titular)	
	
// 	} else {

// 		saque.saldo -= valor
// 		deposito.saldo += valor

// 		println("Transação realizada com sucesso")
// 		println("Valor:", fmt.Sprintf("%.2f", valor))
// 		println("Saque realizado na Conta do cliente:", saque.Titular)
// 		println("Depósito realizado na Conta do cliente:", deposito.Titular)	

// 	}

// 	println("---------------------------------------------")

// }

func (t *ContaCorrente) Pix(deposito *ContaCorrente, valor float64) bool {

	println("---------------------------------------------")

	if valor < 0 {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do pix é negativo")

		return false

	} else if t.Saldo() < valor {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do pix é superior ao saldo da Conta do cliente:", t.Titular.Nome)

		return false
	
	} else {

		t.saldo -= valor
		deposito.saldo += valor

		println("Transação realizada com sucesso")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("Saque realizado na Conta do cliente:", t.Titular.Nome)
		println("Depósito realizado na Conta do cliente:", deposito.Titular.Nome)	

	}

	return true
}

func (t *ContaCorrente) Sacar (valor float64) bool {

	println("---------------------------------------------")

	if valor < 0 {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do saque é negativo")

		return false

	} else if t.Saldo() < valor {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do saque é superior ao saldo da Conta do cliente:", t.Titular.Nome)	

		return false
	
	} else {

		t.saldo -= valor

		println("Transação realizada com sucesso")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("Saque realizado na Conta do cliente:", t.Titular.Nome)

	}

	return true

}

func (t *ContaCorrente) Depositar (valor float64) bool {

	println("---------------------------------------------")

	if valor < 0 {

		println("Transação cancelada")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("O valor do depósito é negativo")

		return false

	} else {

		t.saldo += valor

		println("Transação realizada com sucesso")
		println("Valor:", fmt.Sprintf("%.2f", valor))
		println("Depósito realizado na Conta do cliente:", t.Titular.Nome)

	}

	return true

}


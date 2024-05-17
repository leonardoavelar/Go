package main

import (
	"banco/contas"
	"banco/clientes"
	"banco/pagamentos"
)

func main() {

	println("---------------------------------------------")	
	println("Criando Conta Poupança Elizabeth")

	beth := contas.ContaPoupanca { Titular: clientes.Cliente { "Elizabeth", "12345678901", "Aposentada" }, Agencia: 654, Conta: 8754, Operacao: 13 }
	beth.Depositar(3690.00)

	println("---------------------------------------------")	
	println("Criando Conta Corrente Leonardo")

	leonardo := contas.ContaCorrente { Titular: clientes.Cliente { "Leonardo Avelar", "12345678901", "Engenheiro de Software" }, Agencia: 1234, Conta: 5678 }
	leonardo.Depositar(57020.50)

	println("---------------------------------------------")	
	println("Criando Conta Corrente Fernanda")
	
	fernanda := contas.ContaCorrente { Titular: clientes.Cliente { "Fernanda", "12345678901", "Psicologa" }}
	fernanda.Depositar(23900.89)

	println("---------------------------------------------")	
	println("Criando Conta Corrente Matheus")
	
	matheus := *new(contas.ContaCorrente)
	matheus.Titular = clientes.Cliente { "Matheus", "12345678901", "Estudante" }
	matheus.Depositar(1290.58)

	println("---------------------------------------------")	
	println("Criando Conta Corrente Daniel")
	
	var daniel contas.ContaCorrente = *new(contas.ContaCorrente)
	daniel.Titular = clientes.Cliente {"Daniel", "12345678901", "Estudante"}
	daniel.Agencia = 1455
	daniel.Conta = 4353
	daniel.Depositar(3455.01)

	println("---------------------------------------------")

	beth.Imprimir()
	leonardo.Imprimir()
	fernanda.Imprimir()
	matheus.Imprimir()
	daniel.Imprimir()

	leonardo.Pix(&fernanda, 12000)
	matheus.Pix(&daniel, 3000)
	leonardo.Pix(&matheus, 2000)
	daniel.Pix(&leonardo, -1000)
	matheus.Sacar(5000)
	matheus.Sacar(2000)
	daniel.Depositar(10000)

	pagamentos.EfetuarPagamento(&beth, 1000)
	pagamentos.EfetuarPagamento(&leonardo, 40000)


	println("---------------------------------------------")

	beth.Imprimir()
	leonardo.Imprimir()
	fernanda.Imprimir()
	matheus.Imprimir()
	daniel.Imprimir()
	
	println("---------------------------------------------")

}
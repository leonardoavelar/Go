package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	println("1")
	goto tres

dois:
	println("2")
	goto fim

tres:
	println("3")
	goto dois

fim:
	println("Leonardo Avelar")

	// Imprime as opções de menu
	servico := menuServicos()

	// Trata a opção de menu selecionada
	executar(servico)

	println()

	os.Exit(1)
}

func menuServicos() int {

	println("Menu de serviços:")
	println("0 - Sair")
	println("1 - Abrir Solicitação")
	println("2 - Consultar Solicitação")
	println("3 - Cancelar Solicitação")
	println("")
	print("Digite uma opção: ")

	var servico int
	fmt.Scan(&servico)

	return servico
}

func executar(servico int) {

	switch servico {
	case 0:
		println("Encerrando chat")
	case 1:
		abrirSolicitacao()
	case 2:
		println("Consultando Solicitação")
	case 3:
		println("Cancelando Solicitação")
	default:
		println("Opção inexistente")
	}

}

func abrirSolicitacao() {

	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)

	if resp.StatusCode == http.StatusOK {
		println("Abrindo Solicitação. StatusCode:", resp.StatusCode)
	} else {
		println("Falha na abertura da Solicitação. StatusCode:", resp.StatusCode)
	}

}

package main

import "time"

const wait = 2

func main() {

	var nomes []string = []string{"Leonardo", "Matheus", "Daniel"}
	teste := append(nomes, "Fernanda")

	familia := append(teste, "Antonio", "Flavia", "Beth", "Dalton", "Ester", "Joaquim", "Leandro")

	imprimiSeparacao()

	imprimiLista(nomes)
	imprimiLista(teste)
	imprimiLista(nomes)
	imprimiLista(familia)

	println("Fim do processamento!!")
	imprimiSeparacao()

}

func imprimiLista(lista []string) {

	time.Sleep(wait * time.Second)

	for i, item := range lista {
		println("Posição: ", i, " - Nome: ", item)
	}

	imprimiSeparacao()
}

func imprimiSeparacao() {

	println()
	println("-------------------------------------------")
	println()

}

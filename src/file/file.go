package main

import (
	"bufio"
	"os"
)

func main() {

	println("Realizando leitura de arquivo")

	fileName := "file.txt"
	file := arquivo(fileName)

	bufio.NewReader(file)

	println("Finalizando processo de leitura de arquivo")

}

func arquivo(fileName string) *os.File {

	var file *os.File
	var err error

	file, err = os.Open(fileName)

	if err != nil {

		file, err = os.Create(fileName)

		if err != nil {
			println("Falha ao gerar o arquivo '", fileName, "' - Erro: ", err)
			os.Exit(99)
		} else {
			println("Arquivo '", fileName, "' criado com sucesso")
		}

	} else {
		println("Arquivo '", fileName, "' aberto com sucesso")
	}

	return file

}

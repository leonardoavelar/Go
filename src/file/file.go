package main

import (
	"bufio"
	"os"
	"io"
)

func main() {

	println("Realizando leitura de arquivo")

	fileName := "file.txt"
	file := arquivo(fileName)

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')
		println(line)

		if err == io.EOF {
			break
		}

	}

	println("Finalizando processo de leitura de arquivo")
	
	file.Close()
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

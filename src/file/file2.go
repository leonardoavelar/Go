package main

import (
	"bufio"
	"os"
	"io"
	"net/http"
	"strings"
)

func main() {

	println("Realizando leitura de arquivo")
	println("--------------------------------------------------------")

	sites := "file.txt"
	fileSites := openFile(sites)

	reader := bufio.NewReader(fileSites)

	for {		

		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		
		lineAux := strings.Split(string(line), ";")
		url := lineAux[0]
		fileName  := lineAux[1]

		println(url)
		println(fileName)

		// get content
		content := getContent(url)

		// save content
		saveContent(fileName, content)

		println("--------------------------------------------------------")
		
	}

	println("Finalizando processo de leitura de arquivo")
	
	fileSites.Close()
}

func openFile(fileName string) *os.File {

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

func getContent(url string) string {

	resp, err := http.Get(url)

	if err != nil {

		println("Falha ao acesso o site '", url, "' - Erro: ", err)
		return ""

	} else if resp.StatusCode != http.StatusOK {

		println("Falha ao acesso o site '", url, "' - Status: ", resp.StatusCode)
		return ""

	} else {

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {

			println("Falha ao acesso o site '", url, "' - Erro: ", err)
			return ""
	
		}

		content := string(body)

		return content
	}

}

func saveContent(fileName string, content string) {

	fileNameAux := fileName + ".html"		
	println(fileNameAux)

	err := os.Remove(fileNameAux)

	if os.IsExist(err) {

		println("Falha ao excluir o arquivo '", fileNameAux,"' - Erro: ", err)

	} else {

		println("Arquivo '", fileNameAux,"' excluído com sucesso")

		file, err := os.OpenFile(fileNameAux, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {

			println("Falha ao criar/abrir o arquivo '", fileNameAux,"' - Erro: ", err)

		} else {

			// Armazena o conteudo do site
			file.WriteString(content)
			
			println("Arquivo '", fileNameAux,"' criado com sucesso")

		}

		file.Close()

	}	

}
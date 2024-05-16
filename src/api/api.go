package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	urlApi := "http://localhost:5109/Weather"

	resp, err := http.Get(urlApi)

	if err != nil {
		println("Ocorreu o seguinte erro na request. Erro: ", err)
		os.Exit(99)
	}

	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	content := string(body)

	println(resp.StatusCode)
	println(content)

}

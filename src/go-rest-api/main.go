package main

import (
	"fmt"

	"github.com/leonardoavelar/Go/src/go-rest-api/database"
	"github.com/leonardoavelar/Go/src/go-rest-api/rotas"
)

func main() {

	database.ConectaDB()

	fmt.Print("Inciando servidor Rest com Go")
	rotas.HandleRequest()
}

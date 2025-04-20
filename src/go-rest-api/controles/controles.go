package controles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonardoavelar/Go/src/go-rest-api/database"
	"github.com/leonardoavelar/Go/src/go-rest-api/modelos"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home Page")
}

func GetUsuarios(w http.ResponseWriter, r *http.Request) {

	var usuarios []modelos.Usuario
	database.DB.Find(&usuarios)

	json.NewEncoder(w).Encode(usuarios)
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var usuario modelos.Usuario
	database.DB.First(&usuario, id)

	json.NewEncoder(w).Encode(usuario)
}

func PostUsuario(w http.ResponseWriter, r *http.Request) {

	var usuario modelos.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)
	database.DB.Create(&usuario)

	json.NewEncoder(w).Encode(usuario)
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var usuario modelos.Usuario
	database.DB.Delete(&usuario, id)

	json.NewEncoder(w).Encode(usuario)
}

func PutUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var usuario modelos.Usuario
	database.DB.First(&usuario, id)

	json.NewDecoder(r.Body).Decode(&usuario)
	database.DB.Save(&usuario)

	json.NewEncoder(w).Encode(usuario)
}

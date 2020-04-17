package routes

import (
	"net/http"

	"github.com/luis/crud/controllers"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/add", controllers.Adicionar)
	http.HandleFunc("/inserir", controllers.Inserir)
	http.HandleFunc("/excluir", controllers.Excluir)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/atualizar", controllers.Atualizar)
}

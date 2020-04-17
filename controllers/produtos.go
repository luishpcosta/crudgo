package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/luis/crud/dao"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

/*Pagina principal*/
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := dao.BuscarTodosProduto()
	temp.ExecuteTemplate(w, "Index", produtos)
}

/*pagina de adição de produtos */
func Adicionar(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Adicionar", nil)
}

/*Ação de inserir*/
func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter o preço", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter o preço", err)
		}

		dao.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

/*Ação de excluir*/
func Excluir(w http.ResponseWriter, r *http.Request) {
	idproduto := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idproduto)

	if err != nil {
		log.Println("Erro ao converter o ID", err)
	}

	dao.DeletarProduto(id)
	http.Redirect(w, r, "/", 301)
}

/*Pagina de edição*/
func Editar(w http.ResponseWriter, r *http.Request) {
	idproduto := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idproduto)
	if err != nil {
		log.Println("Erro ao converter o ID", err)
	}
	produto := dao.BuscarProdutoPorId(id)
	temp.ExecuteTemplate(w, "Editar", produto)

}

/*Ação de atualizar*/
func Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter o preço", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter o preço", err)
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro ao converter o preço", err)
			http.Redirect(w, r, "/", 400)
			return
		}

		dao.AtualizaProduto(idConvertido, quantidadeConvertida, nome, descricao, precoConvertido)
	}

	http.Redirect(w, r, "/", 301)

}

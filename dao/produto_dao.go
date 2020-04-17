package dao

import (
	"github.com/luis/crud/db"
	"github.com/luis/crud/models"
)

/*
	Buscar todos os produtos na base de dados
*/
func BuscarTodosProduto() []models.Produto {
	produto := models.Produto{}
	produtos := []models.Produto{}

	db := db.Conectar()
	findAll, err := db.Query("SELECT * FROM  produto")
	if err != nil {
		panic(err.Error())
	}

	for findAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = findAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.ID = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
		produtos = append(produtos, produto)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.Conectar()
	insert, err := db.Prepare("insert into produto(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletarProduto(id int) {
	db := db.Conectar()
	delete, err := db.Prepare("delete from produto where id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}
func BuscarProdutoPorId(id int) models.Produto {
	db := db.Conectar()
	findByid, err := db.Query("SELECT * FROM  produto where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := models.Produto{}
	if findByid.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = findByid.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.ID = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

	}
	defer db.Close()
	return produto

}

func AtualizaProduto(id, quantidade int, nome, descricao string, preco float64) {
	db := db.Conectar()
	update, err := db.Prepare("update produto set quantidade=$1, nome=$2, descricao=$3, preco=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(quantidade, nome, descricao, preco, id)
	defer db.Close()
}

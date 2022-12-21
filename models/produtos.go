package models

import (
	"log"

	"github.com/Gabriel-Newton-dev/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from public.produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	// criei essa variavel p, que irá receber apenas 1 produto, eu irei armanezar o que vem do banco de dados.
	// criei variavel produto para receber o slice do Produto{}

	// aqui tem que ser o mesmo nome da Struct
	p := Produto{}
	produtos := []Produto{}

	// criamos um for para verificar linha a linha, ou seja o selectDeTodosOsProdutos. next, próxima linha.
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64

		// iremos scanear linha a linha, irei guardar em uma variavel de erro, e quero que fique armazenado dentro da memória do meu computador ( &)
		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &valor, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Valor = valor
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		log.Println("Erro na inserção no banco de dados, função inserir", nil)
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()

}

func DeletarProduto(id string) {
	db := db.ConectaComBancoDeDados()

	// delete from(dos) produtos(tabela) where(onde)
	deletarOPrudo, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		log.Println("Erro na função deletar produtos.")
	}

	deletarOPrudo.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Valor = preco
		produtoParaAtualizar.Quantidade = quantidade

	}

	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome string, descricao string, valor float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, valor, quantidade, id)
	defer db.Close()
}

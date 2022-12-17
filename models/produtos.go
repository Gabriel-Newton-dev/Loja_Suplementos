package models

import "github.com/Gabriel-Newton-dev/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from public.produtos")
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

}

package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=loja_suplementos password=G1ogo@2060 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {

	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from public.produtos")
	if err != nil {
		panic(err.Error())
	}

	// criei essa variavel p, que irá receber apenas 1 produto, eu irei armanezar o que vem do banco de dados.
	// criei variavel produto para receber o slice do Produto{}

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

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}

// Produto := []Produtos{
// 	{"Creatina Universal", "A melhor creatina do mercado", 119.89, 10},
// 	{"HGH", "Otimizador de hormonio GH", 350.99, 5},
// 	{"Whey Protein", "O whey protein tem alta concentração de cálcio", 59.99, 10},
// 	{"Albumina", "Clara de ovo", 49.99, 50},
// 	{"Ômega 3", "Multivitaminíco", 32.99, 30},
// 	{"Natubolic", "Suplemento", 91.99, 4},
// }

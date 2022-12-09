package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produtos struct {
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func conectaComBancoDeDados() *sql.DB {
	conexao := "dbname=loja_suplementos password=G1ogo@2060 host=localhost sslmode=disable"
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

	temp.ExecuteTemplate(w, "Index", selectDeTodosOsProdutos)
}

// Produto := []Produtos{
// 	{"Creatina Universal", "A melhor creatina do mercado", 119.89, 10},
// 	{"HGH", "Otimizador de hormonio GH", 350.99, 5},
// 	{"Whey Protein", "O whey protein tem alta concentração de cálcio", 59.99, 10},
// 	{"Albumina", "Clara de ovo", 49.99, 50},
// 	{"Ômega 3", "Multivitaminíco", 32.99, 30},
// 	{"Natubolic", "Suplemento", 91.99, 4},
// }

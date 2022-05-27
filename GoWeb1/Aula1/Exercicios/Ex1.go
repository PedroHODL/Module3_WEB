package Exercicios

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func Ex1() {
	produto := product{
		231543,
		"Cadeira",
		"#182746",
		1500.0,
		9000,
		583930,
		true,
		time.Now(),
	}

	produto2 := product{
		956285,
		"Cadeira2",
		"#182746",
		2000.0,
		5000,
		583931,
		false,
		time.Now().Add(10),
	}

	loja := products{}
	loja.Adicionar(produto)
	loja.Adicionar(produto2)

	jsonData, err := json.Marshal(loja)
	if err != nil {
		panic(err)
	}

	arquivo, err := os.Create("./GoWeb1/Aula1/products.json")
	if err != nil {
		panic(err)
	}

	defer arquivo.Close()

	_, err = arquivo.Write(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo products.json criado com sucesso!")
}

type products []product

type product struct {
	ID            int       `json:"id"`
	Nome          string    `json:"name"`
	Cor           string    `json:"color"`
	Preco         float64   `json:"price"`
	Estoque       int       `json:"stock"`
	Codigo        int       `json:"code"`
	Publicado     bool      `json:"published"`
	DataTransacao time.Time `json:"date"`
}

func (p *products) Adicionar(prod product) {
	*p = append(*p, prod)
}

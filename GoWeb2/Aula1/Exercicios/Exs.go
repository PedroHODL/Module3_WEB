package Exercicios

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Products []Product

type Product struct {
	ID    int     `json:"id" binding:"required"`
	Nome  string  `json:"name" binding:"required"`
	Cor   string  `json:"color" binding:"required"`
	Preco float64 `json:"price" binding:"required"`
}

func (p *Products) Adicionar(prod Product) {
	*p = append(*p, prod)
}

var produtos Products

func Exs() {
	server := gin.Default()

	//Ex1
	server.POST("/ex1", func(c *gin.Context) {
		var produto Product
		err := c.BindJSON(&produto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": err,
			})
			return
		}

		var id int = 1
		if len(produtos) > 0 {
			id = produtos[len(produtos)-1].ID + 1
		}
		produto.ID = id
		produtos = append(produtos, produto)

		c.JSON(http.StatusOK, gin.H{
			"data": produto,
		})
	})

	//Ex2
	server.POST("/ex2", func(c *gin.Context) {
		var p Product
		err := c.ShouldBindJSON(&p)
		if err != nil {
			if p.ID == 0 {
				str := fmt.Sprintf("campo %s é obrigatório", "ID")
				c.JSON(http.StatusNotAcceptable, gin.H{
					"erro": str,
				})
				return
			}

			if p.Nome == "" {
				str := fmt.Sprintf("campo %s é obrigatório", "nome")
				c.JSON(http.StatusNotAcceptable, gin.H{
					"erro": str,
				})
				return
			}

			if p.Cor == "" {
				str := fmt.Sprintf("campo %s é obrigatório", "cor")
				c.JSON(http.StatusNotAcceptable, gin.H{
					"erro": str,
				})
				return
			}

			if p.Preco == 0.0 {
				str := fmt.Sprintf("campo %s é obrigatório", "preço")
				c.JSON(http.StatusNotAcceptable, gin.H{
					"erro": str,
				})
				return
			}
			return
		}

		produtos = append(produtos, p)

		c.JSON(http.StatusOK, gin.H{
			"data": p,
		})
	})

	//Ex3
	server.POST("/ex3", func(c *gin.Context) {
		var produto Product
		err := c.BindJSON(&produto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": err,
			})
			return
		}

		token := c.GetHeader("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"erro": "você não tem permissão para fazer a solicitação solicitada",
			})
			return
		}

		var id int = 1
		if len(produtos) > 0 {
			id = produtos[len(produtos)-1].ID + 1
		}
		produto.ID = id
		produtos = append(produtos, produto)

		c.JSON(http.StatusOK, gin.H{
			"data": produto,
		})
	})

	server.Run()
}

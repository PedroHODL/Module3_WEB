package Exercicios

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Exs() {
	server := gin.Default()
	//Ex1
	server.GET("/hello", HelloPage)

	//Ex2
	//server.GET("/empregado/:id", SearchEmpregado)
	server.GET("/temas/:id", searchTemas)
	server.Run()
}

func HelloPage(c *gin.Context) {
	c.String(http.StatusOK, "Bem vindo %s!\nPágina = %s", c.Query("name"), c.Query("page"))
}

var temas = map[string]string{
	"1": "Ação",
	"2": "Aventura",
	"3": "Romance",
	"4": "Terror",
	"5": "Suspense",
}

func searchTemas(c *gin.Context) {
	temas, ok := temas[c.Param("id")]
	if ok {
		c.String(http.StatusOK, "Informação do tema %s = %s", c.Param("id"), temas)
	} else {
		c.String(http.StatusNotFound, "Informação do tema não encontrado!")
	}
}

// Exemplo da aula

/**
var empregado = map[string]string{
	"123": "Empregado A",
	"456": "Empregado B",
	"789": "Empregado C",
}

func SearchEmpregado(c *gin.Context) {
	empregado, ok := empregado[c.Param("id")]
	if ok {
		c.String(http.StatusOK, "Informação do empregado %s, nome %s", c.Param("id"), empregado)
	} else {
		c.String(http.StatusNotFound, "Informação de Empregado não encontrado!")
	}
}
**/

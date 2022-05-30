package Exercicios

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ex1() {
	server := gin.Default()
	server.GET("/", RootPage)
	//Ex1
	server.GET("/empregado/:id", SearchEmpregado)

	//Ex2
	//server.GET("/temas/:id", searchTemas)
	server.Run()
}

var empregado = map[string]string{
	"123": "Empregado A",
	"456": "Empregado B",
	"789": "Empregado C",
}

func RootPage(c *gin.Context) {
	c.String(http.StatusOK, "Bem vindo!")
}

func SearchEmpregado(c *gin.Context) {
	empregado, ok := empregado[c.Param("id")]
	if ok {
		c.String(http.StatusOK, "Informação do empregado %s, nome %s", c.Param("id"), empregado)
	} else {
		c.String(http.StatusNotFound, "Informação de Empregado não encontrado!")
	}
}

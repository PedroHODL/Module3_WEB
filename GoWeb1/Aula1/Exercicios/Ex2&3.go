package Exercicios

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Ex2_3() {
	gin.SetMode("release")
	router := gin.Default()

	arq, err := os.ReadFile("./GoWeb1/Aula1/message.json")
	if err != nil {
		panic(err)
	}

	//Ex2
	router.GET("/", func(c *gin.Context) {
		c.Data(200, "json", arq)
	})

	//Ex3
	group := router.Group("/products")
	{
		group.GET("/all", GetAll)
	}

	router.Run()
}

func GetAll(c *gin.Context) {
	arq, err := os.ReadFile("./GoWeb1/Aula1/products.json")
	if err != nil {
		panic(err)
	}
	c.Data(200, "json", arq)
}

package Exercicios

import (
	"github.com/gin-gonic/gin"
)

func Ex2_3() {
	gin.SetMode("release")
	router := gin.Default()

	//Ex2
	router.GET("/", func(c *gin.Context) {
		c.File("./GoWeb1/Aula1/message.json")
	})

	//Ex3
	group := router.Group("/products")
	{
		group.GET("/all", GetAll)
	}

	router.Run()
}

func GetAll(c *gin.Context) {
	c.File("./GoWeb1/Aula1/products.json")
}

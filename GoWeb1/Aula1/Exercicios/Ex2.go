package Exercicios

import (
	"github.com/gin-gonic/gin"
	"os"
)

func IniciarGin(c *gin.Context, jsonData []byte) {
	c.JSON(200, jsonData)
}

func Ex2() {
	gin.SetMode("release")
	router := gin.Default()

	arq, err := os.ReadFile("./GoWeb1/Aula1/message.json")
	if err != nil {
		panic(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.Data(200, "json", arq)
	})

	router.Run()
}

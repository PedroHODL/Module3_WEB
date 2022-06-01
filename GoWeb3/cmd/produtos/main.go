package main

import (
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/cmd/produtos/handler"
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/internal/produto"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode("release")
	_ = godotenv.Load("./.env")

	rep := produto.NewRepository()
	service := produto.NewService(rep)
	p := handler.NewProduct(service)

	server := gin.Default()

	r := server.Group("/produtos")
	r.POST("/", p.CreateProduct)
	r.GET("/", p.GetAll)
	r.PUT("/:id", p.Update)
	r.PATCH("/:id", p.UpdateName)
	r.DELETE("/:id", p.DeleteProduct)

	server.Run()
}

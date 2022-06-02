package main

import (
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/cmd/produtos/handler"
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/internal/produto"
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode("release")
	_ = godotenv.Load("./.env")

	db := store.New(store.FileType, "./produtos.json")

	rep := produto.NewRepository(db)
	service := produto.NewService(rep)
	p := handler.NewProduct(service)

	server := gin.Default()

	r := server.Group("/produtos")
	{
		r.Use(p.TokenAuthMiddleware())

		r.POST("/", p.CreateProduct)
		r.GET("/", p.GetAll)
		r.PUT("/:id", p.IdVerificatorMiddleware, p.Update)
		r.PATCH("/:id", p.IdVerificatorMiddleware, p.UpdateName)
		r.DELETE("/:id", p.IdVerificatorMiddleware, p.DeleteProduct)
	}

	server.Run()
}

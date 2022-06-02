package main

import (
	"os"

	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/cmd/produtos/handler"
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/docs"
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/internal/produto"
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.br/pt_br/termos-e-condicoes

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.br/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	gin.SetMode("release")
	_ = godotenv.Load("./.env")

	db := store.New(store.FileType, "./produtos.json")

	rep := produto.NewRepository(db)
	service := produto.NewService(rep)
	p := handler.NewProduct(service)

	server := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r := server.Group("/produtos")
	{
		r.Use(p.TokenAuthMiddleware)

		r.POST("/", p.CreateProduct)
		r.GET("/", p.GetAll)
		r.PUT("/:id", p.IdVerificatorMiddleware, p.Update)
		r.PATCH("/:id", p.IdVerificatorMiddleware, p.UpdateName)
		r.DELETE("/:id", p.IdVerificatorMiddleware, p.DeleteProduct)
	}

	server.Run()
}

package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/internal/produto"
	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name        string  `json:"name"`
	ProductType string  `json:"type"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
}

type Product struct {
	service produto.Services
}

func NewProduct(p produto.Services) Product {
	new := Product{p}
	return new
}

func (p *Product) TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("Variavel de sistema TOKEN vazia")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(web.DecodeError(http.StatusUnauthorized, "token vazio"))
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(web.DecodeError(http.StatusUnauthorized, "token inválido"))
			return
		}

		ctx.Next()
	}
}

func (p *Product) IdVerificatorMiddleware(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "id não é alphanumérico"))
		return
	}

	if 0 > id && id > p.service.LastID() {
		ctx.AbortWithStatusJSON(web.DecodeError(http.StatusBadRequest, "id fora do limite"))
		return
	}

	ctx.Next()
}

// Listar todos os produtos godoc
// @Summary Listar todos os produtos
// @Tags Products
// @Description Lista todos os produtos no banco de dados do Meli
// @Accept  json
// @Produce  json
// @Param token header string true "Token"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /produtos [get]
func (p *Product) GetAll(ctx *gin.Context) {
	prod, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(web.NewResponse(http.StatusOK, prod))
	return
}

// Adicionar um novo produto godoc
// @Summary Adicionar um novo produto
// @Tags Products
// @Description Adiciona um novo produto ao banco de dados do Meli
// @Accept  json
// @Produce  json
// @Param token header string true "Token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Router /produtos [post]
func (p *Product) CreateProduct(ctx *gin.Context) {
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'nome' é obrigatório"))
		return
	}

	if req.ProductType == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'type' é obrigatório"))
		return
	}

	if req.Count <= 0 {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'count' não pode ser menor que 1"))
		return
	}

	if req.Price < 0 {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'price' não pode ser menor que 0"))
		return
	}

	prod, err := p.service.Create(req.Name, req.ProductType, req.Count, req.Price)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, prod))
}

// Substituir um produto godoc
// @Summary Substituir um produto
// @Tags Products
// @Description Atualizar um produto que já esta no banco de dados do Meli
// @Accept  json
// @Produce  json
// @Param 	token 	header 	string 	true "Token"
// @Param 	id 		path 	int  	true  "Product ID"
// @Param 	product body 	request true "Product to Update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /produtos/{id} [put]
func (p *Product) Update(ctx *gin.Context) {
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'nome' é obrigatório"))
		return
	}

	if req.ProductType == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'type' é obrigatório"))
		return
	}

	if req.Count <= 0 {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'count' não pode ser menor que 1"))
		return
	}

	if req.Price < 0 {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'price' não pode ser menor que 0"))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	prod, err := p.service.Update(id, req.Name, req.ProductType, req.Count, req.Price)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, prod))
	return
}

// Trocar o nome de um produto godoc
// @Summary Trocar o nome de um produto
// @Tags Products
// @Description Atualizar o nome de um produto existente no banco de dados do Meli
// @Accept  json
// @Produce  json
// @Param 	token 	header 	string 	true "Token"
// @Param 	id 		path 	int  	true  "Product ID"
// @Param 	product body 	request true "Product to Patch"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /produtos/{id} [patch]
func (p *Product) UpdateName(ctx *gin.Context) {
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "campo 'nome' é obrigatório"))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	prod, err := p.service.UpdateName(id, req.Name)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, prod))
	return
}

// Deletar um produto godoc
// @Summary Deletar um produto
// @Tags Products
// @Description Deletar um produto do banco de dados do Meli
// @Accept  json
// @Produce  json
// @Param 	token 	header 	string 	true "Token"
// @Param 	id 		path 	int  	true  "Product ID to Delete"
// @Success 200 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /produtos/{id} [delete]
func (p *Product) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := p.service.DeleteProduct(id)
	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, err.Error()))
		return
	}

	prod := fmt.Sprintf("O produto %d foi removido", id)
	ctx.JSON(web.NewResponse(http.StatusOK, prod))

}

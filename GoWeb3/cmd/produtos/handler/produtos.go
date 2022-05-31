package handler

import (
	"net/http"

	"github.com/PedroHODL/Module3_WEB.git/GoWeb3/internal/produto"
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

func (p *Product) GetAll(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": "token inválido"})
		return
	}

	prod, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}
	ctx.JSON(200, prod)
}

func (p *Product) CreateProduct(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": "token inválido"})
		return
	}

	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	prod, err := p.service.Create(req.Name, req.ProductType, req.Count, req.Price)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}
	ctx.JSON(200, prod)
}

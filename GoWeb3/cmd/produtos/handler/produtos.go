package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": "token inválido"})
		return
	}

	prod, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}
	ctx.JSON(200, prod)
	return
}

func (p *Product) CreateProduct(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
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

func (p *Product) Update(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": "token inválido"})
		return
	}

	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Aqui"})
		return
	}

	prod, err := p.service.Update(id, req.Name, req.ProductType, req.Count, req.Price)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(200, prod)
	return
}

func (p *Product) UpdateName(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": "token inválido"})
		return
	}

	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "O nome é obrigatório"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	prod, err := p.service.UpdateName(id, req.Name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(200, prod)
	return
}

func (p *Product) DeleteProduct(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": "token inválido"})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if id > p.service.LastID() {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID out of bounds"})
		return
	}

	err = p.service.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})

}

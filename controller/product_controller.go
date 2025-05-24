package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.ProductUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	idParam := ctx.Param("productId")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "ID do produto precisa ser um numero",
		})
		return
	}

	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Dados invalidos",
		})
		return
	}

	// garantir que o id do path corresponde ao produto
	product.ID = productID

	// verf existencia
	existingProduct, err := p.ProductUseCase.GetProductById(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Erro ao verificar produto",
		})
		return
	}

	if existingProduct == nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Produto nao encontrado",
		})
		return
	}

	//atualizar
	if err := p.ProductUseCase.UpdateProduct(product); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Erro ao atualizar produto",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("productId")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "ID do produto precisa ser um numero",
		})
		return
	}

	existingProduct, err := p.ProductUseCase.GetProductById(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Erro ao verificar produto",
		})
		return
	}

	if existingProduct == nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Produto nao encontrado",
		})
		return
	}

	// deletar
	if err := p.ProductUseCase.DeleteProduct(productID); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Erro ao deletar o produto",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Produto deletado ccom sucesso",
	})
}

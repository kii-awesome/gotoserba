package productcontroller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kii-awesome/gotoserba/data"
	"github.com/kii-awesome/gotoserba/models"
	"gorm.io/gorm"
)

var db = data.ConnectDb()

func GetAllProduct(ctx *gin.Context) {
	var products []models.Product

	result := db.Find(&products)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   result.Error.Error(),
			"message": "Failed for getting product",
		})
		return
	}

	response := &models.ProductWebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   products,
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateProduct(ctx *gin.Context) {
	var request *models.Product

	err := ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
			"message": "should with json format",
		})
	}

	result := db.Create(&models.Product{
		Name:        request.Name,
		Description: request.Description,
		Type:        request.Type,
		Price:       request.Price,
		Stock:       request.Stock,
		Image:       request.Image,
		Category:    request.Category,
	})
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   result.Error.Error(),
			"message": "failed to create product",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "succes create product",
	})
}

func GetProductByName(ctx *gin.Context) {
	name := ctx.Param("name")

	product := models.Product{}

	err := db.Where(&models.Product{Name: name}).Take(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   errors.Is(err, gorm.ErrRecordNotFound),
			"message": "product not found",
		})
		return
	}

	response := &models.Product{
		Name:        product.Name,
		Description: product.Description,
		Type:        product.Type,
		Price:       product.Price,
		Stock:       product.Stock,
		Image:       product.Image,
		Category:    product.Category,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func EditProduct(ctx *gin.Context) {
	var request models.Product
	name := ctx.Param("name")

	err := ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"code":    http.StatusBadRequest,
			"message": "fail for json request",
		})
	}

	err = db.Where("name = ?", name).Updates(models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		Image:       request.Image,
	}).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"code":    http.StatusAccepted,
		"message": "succes updated product",
	})
}

func DeleteProduct(ctx *gin.Context)  {
	var request models.Product

	err := ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"code":    http.StatusBadRequest,
			"message": "product not found",
		})
		return
	}

	err = db.Where("name = ?", request.Name).Delete(models.Product{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   errors.Is(err, gorm.ErrRecordNotFound),
			"message": "failed to delete product",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"code":    http.StatusAccepted,
		"message": "succes delete product",
	})
}
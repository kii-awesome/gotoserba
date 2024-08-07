package productmodel

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kii-awesome/gotoserba/data"
	"github.com/kii-awesome/gotoserba/entities"
	"gorm.io/gorm"
)

var db = data.ConnectDb()

func GetAllProduct(ctx *gin.Context) {
	var products []entities.Product

	result := db.Find(&products)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": result.Error.Error(),
			"message": "Failed for getting product",
		})
	}

	response := &entities.ProductWebResponse{
		Code: http.StatusOK,
		Status: "success",
		Data: products,
	} 

	ctx.JSON(http.StatusOK, response)
}

func CreateProduct(ctx *gin.Context)  {
	var data *entities.ProductRequest

	err := ctx.ShouldBindBodyWithJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": err.Error(),
			"message": "should with json format",
		})
	}

	product := &entities.Product{
		Name : data.Name,
		Description : data.Description,
		Type : data.Type,
		Price : data.Price,
		Stock : data.Stock,
		Image : data.Image,
		Category : data.Category,
		CreatedAt : time.Now(),
	}

	result := db.Create(&product)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": result.Error.Error(),
			"message": "failed to create product",
		})
		return
	}

	response := &entities.ProductResponse{
		Name: product.Name,
		Description: product.Description,
		Type: product.Type,
		Price: product.Price,
		Stock: product.Stock,
		Image: product.Image,
		Category: product.Category,
		CreatedAt: product.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, response)
}

func GetProductByName(ctx *gin.Context)  {
	name :=  ctx.Param("name")
	product := entities.Product{}

	err := db.Where(&entities.Product{Name:name}).Take(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"error": errors.Is(err, gorm.ErrRecordNotFound),
			"message": "product not found",
		})
		return
	}

	response := &entities.ProductResponse{
		Name: product.Name,
		Description: product.Description,
		Type: product.Type,
		Price: product.Price,
		Stock: product.Stock,
		Image: product.Image,
		Category: product.Category,
		CreatedAt: product.CreatedAt,
	}
	
	ctx.JSON(http.StatusOK, response)
}
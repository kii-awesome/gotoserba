package categorymodel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kii-awesome/gotoserba/data"
	"github.com/kii-awesome/gotoserba/entities"
)

var db = data.ConnectDb()

func GetAllCategory(ctx *gin.Context) {
	var categories []entities.Category
	result := db.Find(&categories)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": result.Error.Error(),
			"message": "Failed for getting category",
		})
	}
	
	response := &entities.CategoryWebResponse{
        Code:   http.StatusOK,
        Status: "success",
        Data:   categories,
    }

    ctx.JSON(http.StatusOK, response)
}

func CreateCategory(ctx *gin.Context)  {
	var data *entities.CategoryRequest

	// Mengikat badan permintaan json ke struktur badan permintaan
	err := ctx.ShouldBindBodyWithJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": err.Error(),
			"message": "Failed create category",
		})
	}

	category := entities.Category{}
	category.Name = data.Name

	tx := db.Create(&category)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": tx.Error,
			"code": http.StatusBadRequest,
			"message": "failed to create category",
		})
		return
	}
	
	var response entities.CategoryResponse
	response.ID = uint(category.ID)
	response.Name = category.Name

	ctx.JSON(http.StatusCreated, response)
}
package categorycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kii-awesome/gotoserba/data"
	"github.com/kii-awesome/gotoserba/models"
)

var db = data.ConnectDb()

func GetAllCategory(ctx *gin.Context) {
	var categories []models.Category

	result := db.Find(&categories)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": result.Error.Error(),
			"message": "Failed for getting category",
		})
	}
	
	response := &models.CategoryWebResponse{
        Code:   http.StatusOK,
        Status: "success",
        Data:   categories,
    }

    ctx.JSON(http.StatusOK, response)
}

func CreateCategory(ctx *gin.Context)  {
	var request models.Category

	err := ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": err.Error(),
			"message": "should with json format",
		})
	}

	result := db.Create(&models.Category{
		Name:        request.Name,
		Description: request.Description,
	})
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
			"code": http.StatusBadRequest,
			"message": "failed to create category",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "succes create category",
	})
}
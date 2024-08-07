package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kii-awesome/gotoserba/models/categorymodel"
	"github.com/kii-awesome/gotoserba/models/productmodel"
)

func Routes() {
	gin.ForceConsoleColor()
	
	r := gin.Default()

	// product
	r.GET("/api/product", productmodel.GetAllProduct)
	r.POST("/api/product", productmodel.CreateProduct)
	r.GET("/api/product/:name", productmodel.GetProductByName)
	
	// category
	r.GET("/api/category", categorymodel.GetAllCategory)
	r.POST("/api/category", categorymodel.CreateCategory)
	
	err := r.Run("localhost:3000")
	if err != nil {
		panic(err)
	}
}

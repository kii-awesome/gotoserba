package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kii-awesome/gotoserba/controller/categorycontroller"
	"github.com/kii-awesome/gotoserba/controller/productcontroller"
)

func Routes() {
	gin.ForceConsoleColor()

	r := gin.Default()

	// product
	r.GET("/api/product", productcontroller.GetAllProduct)
	r.POST("/api/product", productcontroller.CreateProduct)
	r.GET("/api/product/:name", productcontroller.GetProductByName)
	r.PUT("/api/product/edit/:name", productcontroller.EditProduct)
	r.DELETE("api/product", productcontroller.DeleteProduct)
	// category
	r.GET("/api/category", categorycontroller.GetAllCategory)
	r.POST("/api/category", categorycontroller.CreateCategory)

	err := r.Run("localhost:3000")
	if err != nil {
		panic(err)
	}
}

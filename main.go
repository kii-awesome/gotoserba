package main

import (
	_ "github.com/gin-gonic/gin"
	"github.com/kii-awesome/gotoserba/data"
	"github.com/kii-awesome/gotoserba/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = data.ConnectDb()
)

func main() {
	defer data.DiscconectDb(db)

	routes.Routes()
}

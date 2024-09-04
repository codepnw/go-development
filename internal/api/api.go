package api

import (
	"fmt"
	"os"

	"github.com/codepnw/godevelopment/internal/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewRoutes(version string) {
	r := gin.Default()
	port := os.Getenv("APP_PORT")

	db, err := gorm.Open(sqlite.Open(models.SessionFile()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&models.Session{}); err != nil {
		panic("unable to mae migrations sessions tables")
	}

	r.Use(cors.Default())

	// routes
	AuthRoutes(r, version)
	UserRoutes(r, version)

	fmt.Printf("listening at port %s", port)
	r.Run(":" + port)
}

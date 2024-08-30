package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func NewRoutes(version string) {
	r := gin.Default()
	r.Use(cors.Default())

	AuthRoutes(r, version)
	UserRoutes(r, version)

	fmt.Println("listening at port 8080")
	r.Run(":8080")
}

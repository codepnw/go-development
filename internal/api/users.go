package api

import (
	"github.com/codepnw/godevelopment/internal/handlers/users"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, version string) {
	g := r.Group(version + "/users")

	g.POST("/", users.CreateHandler)
	g.GET("/", users.GetHandler)
	g.GET("/:id", users.GetByIdHandler)
	g.PATCH("/:id", users.UpdateHandler)
	g.DELETE("/:id", users.DeleteHandler)
}

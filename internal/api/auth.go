package api

import (
	"github.com/codepnw/godevelopment/internal/handlers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, version string) {
	g := r.Group(version + "/auth")

	g.POST("/login", auth.LoginHandler)
	g.GET("/logout", auth.LogoutHandler)
}

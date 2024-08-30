package api

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRoutes(version string) {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	port := os.Getenv("APP_PORT")
	nums := []string{"auth", "account"}

	r.Use(cors.Default())
	r.Use(sessions.Sessions("golang-dev", store))
	r.Use(sessions.SessionsMany(nums, store))

	// routes
	AuthRoutes(r, version)
	UserRoutes(r, version)

	fmt.Printf("listening at port %s", port)
	r.Run(":" + port)
}

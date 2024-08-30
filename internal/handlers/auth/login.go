package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/codepnw/godevelopment/internal/types/requests"
	"github.com/codepnw/godevelopment/internal/utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user models.User
	var req requests.LoginRequest
	secret := os.Getenv("JWT_SECRET")

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FormatError(c, err)
		return
	}

	user.Email = req.Email

	db := db.GetClientGorm()
	if err := user.GetByAttr(db).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	fmt.Println("the hashed password: ", hash)
	if err := utils.VerifyPassword(hash, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	if token, err := utils.GenerateJWT(user, secret); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
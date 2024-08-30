package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/codepnw/godevelopment/internal/utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var req models.User
	secret := os.Getenv("JWT_SECRET")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "invalid payload",
		})
		return
	}

	db := db.GetClientGorm()
	if err := req.GetByAttr(db).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	password := "mypassword"
	hash, err := utils.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	fmt.Println("the hashed password: ", hash)
	if err := utils.VerifyPassword(hash, password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	if token, err := utils.GenerateJWT(req, secret); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
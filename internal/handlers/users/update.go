package users

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/codepnw/godevelopment/internal/utils"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(c *gin.Context) {
	var req models.User
	var access_token string
	var err error
	secret := os.Getenv("JWT_SECRET")

	if access_token, err = utils.GetToken(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token or missing token"})
		return
	}

	claims, err := utils.ValidateJWT(access_token, secret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	fmt.Println(claims)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"reason": "invalid payload provided",
		})
		return
	}

	db := db.GetClientGorm()
	if err := req.Update(db, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, req)
}

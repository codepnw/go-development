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

func DeleteHandler(c *gin.Context) {
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

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request missing user id",
		})
		return
	}

	req.ID = id
	db := db.GetClientGorm()
	if err := req.Delete(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "user deleted"})
}

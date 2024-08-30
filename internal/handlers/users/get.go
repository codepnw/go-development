package users

import (
	"net/http"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	var req models.User

	db := db.GetClientGorm()
	var users []models.User

	if err := req.GetAll(db, users).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success", 
		"users": users,
	})
}

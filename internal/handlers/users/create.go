package users

import (
	"net/http"
	"time"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateHandler(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "invalid payload provided",
		})
		return
	}
	
	req.ID = uuid.New().String()
	req.CreatedAt = time.Now().Local().String()

	db := db.GetClientGorm()
	if err := req.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}
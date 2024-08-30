package users

import (
	"net/http"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/gin-gonic/gin"
)

func GetByIdHandler(c *gin.Context) {
	var req models.User

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request missing user id",
		})
		return
	}

	req.ID = id
	db := db.GetClientGorm()
	if err := req.GetById(db).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, req)
}

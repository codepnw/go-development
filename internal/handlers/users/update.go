package users

import (
	"net/http"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(c *gin.Context) {
	var req models.User

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

package users

import (
	"net/http"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(c *gin.Context) {
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
	if err := req.Delete(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "user deleted"})
}

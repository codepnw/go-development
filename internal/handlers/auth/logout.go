package auth

import (
	"encoding/json"
	"net/http"

	"github.com/codepnw/godevelopment/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LogoutHandler(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		panic("cookie missing in request")
	}

	var session models.Session
	db, err := gorm.Open(sqlite.Open(models.SessionFile()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Where("id = ?", sessionID).First(&session)

	var data map[string]string
	json.Unmarshal(session.Data, &data)

	c.JSON(http.StatusOK, gin.H{"session": session})
}

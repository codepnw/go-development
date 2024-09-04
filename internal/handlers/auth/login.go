package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/codepnw/godevelopment/internal/config/db"
	"github.com/codepnw/godevelopment/internal/models"
	"github.com/codepnw/godevelopment/internal/types/requests"
	"github.com/codepnw/godevelopment/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	if err := utils.VerifyPassword(hash, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	data := map[string]string{"email": user.Email, "user_id": user.ID}
	data_bytes, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Session
	session := models.Session{
		ID:     uuid.NewString(),
		Data:   data_bytes,
		Expiry: time.Now().Add(24 * time.Hour).Unix(),
	}

	gormDB, err := gorm.Open(sqlite.Open("session.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := gormDB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if token, err := utils.GenerateJWT(user, secret); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
	} else {
		// Cookie
		utils.PersistCookie(c, session.ID, token)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

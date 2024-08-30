package auth

import (
	"net/http"

	"github.com/codepnw/godevelopment/internal/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	var sessionKey = "user-role"

	utils.RemoveCookie(c)

	session := sessions.Default(c)
	if session.Get(sessionKey) == nil {
		session.Set(sessionKey, "ADMIN")
		session.Save()
	}

	c.JSON(http.StatusOK, gin.H{"session": ""})
}
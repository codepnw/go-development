package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 86400 = 1 Day (60*60*24)
const cookieAge = 86400

func PersistCookie(c *gin.Context, session_id, token string) {
	c.SetCookie("access_token", token, cookieAge, "/", "localhost", false, true)
	c.SetCookie("cookie_exp", strconv.Itoa(cookieAge), cookieAge, "/", "localhost", false, true)
	c.SetCookie("session_id", session_id, 3600, "/", "", false, true)
	c.SetSameSite(http.SameSiteStrictMode)
}

func RemoveCookie(c *gin.Context) {
	c.SetCookie("access_token", "", cookieAge, "/", "localhost", false, true)
	c.SetCookie("cookie_exp", "", cookieAge, "/", "localhost", false, true)
	c.SetCookie("session_id", "", cookieAge, "/", "", false, true)
	c.SetSameSite(http.SameSiteStrictMode)
}

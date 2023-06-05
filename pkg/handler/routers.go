package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const userKey = "user"

var secret = []byte("secret")

func Engine() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore(secret)
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/login", ShowLoginPage)
	r.POST("/login", Login)
	r.GET("/logout", Logout)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", Me)
		private.GET("/status", Status)
	}
	return r
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userKey)
	if user == nil {
		c.Redirect(http.StatusFound, "/")
		c.Abort()
		return
	}
	c.Next()
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(userKey)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userKey)
	c.HTML(http.StatusOK, "user.html", gin.H{"user": user})
}

func Status(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Parameters can't be empty"})
		return
	}

	if username != "def" || password != "123" {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"message": "Authentication failed"})
		return
	}

	session.Set(userKey, username)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusFound, "/private/me")
}

func Index(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(userKey) != nil {
		c.Redirect(http.StatusFound, "/private/me")
		return
	}

	c.HTML(http.StatusOK, "login.html", nil)
}

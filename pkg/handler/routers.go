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

func engine() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore(secret)
	r.Use(sessions.Sessions("mysession", store))

	r.Static("/static", "./static")

	r.GET("/login", showLoginPage)
	r.POST("/login", login)
	r.GET("/logout", logout)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}
	return r
}

func showLoginPage(c *gin.Context) {
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

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(userKey)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userKey)
	c.HTML(http.StatusOK, "user.html", gin.H{"user": user})
}

func status(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}

func login(c *gin.Context) {
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

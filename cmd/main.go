package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var secret = []byte("secret")

const userKey = "user"

func main() {
	r := engine()

	r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/*")

	r.GET("/", index)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func index(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(userKey) != nil {
		c.Redirect(http.StatusFound, "/private/me")
		return
	}

	c.HTML(http.StatusOK, "login.html", nil)
}

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

package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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

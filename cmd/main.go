package main

import (
	"github.com/drew3k/site/pkg/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := handler.Engine()

	r.Use(gin.Logger())

	r.LoadHTMLGlob("site/templates/*.html")

	r.GET("/", handler.Index)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

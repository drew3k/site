package main

import (
	"github.com/drew3k/site/pkg/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"myproj/site/database"
)

func main() {
	db, err := database.SetupDB()
	if err != nil {
		log.Fatal("Failed to setup database:", err)
		return
	}
	defer db.Close()

	r := handler.Engine()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("site/templates/*.html")

	r.GET("/", handler.Index)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

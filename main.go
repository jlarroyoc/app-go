package main

import (
	"log"
	"info-go/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "./static/css")
	r.Static("/images", "./static/images")

	r.LoadHTMLGlob("templates/*")
	controller.Router(r)

	log.Println("****************Server started...")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package main

import (
	"web/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/articles/:id", controllers.GetArticle)
	router.POST("/articles", controllers.CreateArticle)
	router.DELETE("/articles/:id", controllers.DeleteArticle)

	router.Run(":8081")
}

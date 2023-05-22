package main

import (
	"web/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/categories", controllers.GetCategory)
	router.GET("/articles", controllers.GetArticle)

	return router
}

package utilities

import (
	"web/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/articles/:id", controllers.GetArticle)
	router.POST("/articles", controllers.CreateArticle)
	router.DELETE("/articles/:id", controllers.DeleteArticle)

	return router
}

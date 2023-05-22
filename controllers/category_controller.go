package controllers

import (
	"net/http"
	"web/entities"

	"github.com/gin-gonic/gin"
)

func GetCategory(context *gin.Context) {
	var request entities.IdRequestEntity
	context.BindQuery(request)
	context.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

package controllers

import (
	"net/http"
	"strconv"
	"web/entities/articles"
	repo "web/repositories/articles"
	svc "web/services/articles"

	"github.com/gin-gonic/gin"
)

func GetArticle(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 32)
	if err != nil {
		return
	}

	var article repo.Article = svc.GetArticle(id)

	context.JSON(http.StatusOK, gin.H{
		"title":   article.Title,
		"content": article.Content,
	})
}

func CreateArticle(context *gin.Context) {
	var request articles.ArticleCreateRequestEntity
	context.BindJSON(&request)
	var id = svc.CreateArticle(request)
	context.JSON(http.StatusOK, gin.H{
		"Id": id,
	})
}

func DeleteArticle(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 32)
	if err != nil {
		return
	}
	svc.DeleteArticle(id)
}

package articles

import (
	"web/entities/articles"
	repo "web/repositories/articles"
)

func GetArticle(id int64) repo.Article {
	article := repo.GetArticle(id)
	return article
}

func CreateArticle(request articles.ArticleCreateRequestEntity) int64 {
	id := repo.CreateArticle(repo.Article{Title: request.Title, Content: request.Content})
	return id
}

func DeleteArticle(id int64) {
	repo.DeleteArticle(id)
}

package repostiories

import (
	"time"
	db "web/databases"

	"gorm.io/gorm"
)

type Article struct {
	Id              int64 `gorm:"primaryKey"`
	Title           string
	Content         string
	CreatedDatetime string
	UpdatedDatetime string
	Deleted_At      gorm.DeletedAt
}

func GetArticle(id int64) Article {

	var conn = db.GetDatabaseConnection()
	var article = Article{Id: id}
	conn.First(&article)
	return article
}

func CreateArticle(article Article) int64 {
	var conn = db.GetDatabaseConnection()
	conn.Create(&article)
	return article.Id
}

func DeleteArticle(id int64) {
	var conn = db.GetDatabaseConnection()
	var article = Article{Id: id}
	conn.First(&article)
	conn.Delete(&article)
}

func (Article) TableName() string {
	return "article"
}

func (article *Article) BeforeCreate(db *gorm.DB) (err error) {
	article.CreatedDatetime = time.Now().String()
	article.UpdatedDatetime = time.Now().String()
	return
}

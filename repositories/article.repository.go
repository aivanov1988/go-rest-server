package repositories

import (
	"github.com/aivanov1988/go-rest-server/models"
)

func ReadAllArticles() ([]models.Article, error) {
	var articles []models.Article
	err := db.Model(&articles).Table("articles").Select()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

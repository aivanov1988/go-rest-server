package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aivanov1988/go-rest-server/responses"
)

func GetArticlesList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, responses.ArticleListResponse{{Title: "Article 1"}, {Title: "Article 2"}})
}

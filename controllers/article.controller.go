package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aivanov1988/go-rest-server/helpers"
	"github.com/aivanov1988/go-rest-server/repositories"
	"github.com/aivanov1988/go-rest-server/responses"
)

func GetArticlesList(c *gin.Context) {
	articleList, err := repositories.ReadAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.PrepareErrorResponse(err))
		return
	}
	res, err := helpers.TypeConverter[responses.ArticleListResponse](articleList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.PrepareErrorResponse(err))
		return
	}
	//resStr := helpers.JsonStringify(res)
	c.JSON(http.StatusOK, res)
}

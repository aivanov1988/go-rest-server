package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/aivanov1988/go-rest-server/controllers"
)

func ArticlesRouter(router *gin.Engine) {
	router.GET("/articles", controllers.GetArticlesList)

}

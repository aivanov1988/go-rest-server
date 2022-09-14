package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/aivanov1988/go-rest-server/controllers"
)

func AppRouter(router *gin.Engine) {
	router.GET("/version", controllers.GetAppVersion)
}

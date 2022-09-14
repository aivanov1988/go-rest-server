package routers

import (
	"github.com/gin-gonic/gin"
)

func ProtectedRouter(router *gin.Engine) {
	ArticlesRouter(router)

}

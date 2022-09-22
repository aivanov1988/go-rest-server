package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aivanov1988/go-rest-server/responses"
)

func GetAppVersion(c *gin.Context) {
	c.JSON(http.StatusOK, responses.VersionResponse{Version: "0.0.1"})
}

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/aivanov1988/go-rest-server/config"
	"github.com/aivanov1988/go-rest-server/repositories"
	"github.com/aivanov1988/go-rest-server/routers"
)

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}

	repositories.ConnectToDB(config.Config.DbConf)
	defer repositories.CloseConnection()

	app := gin.Default()
	routers.AppRouter(app)
	//router.Routes(UploadRouter)
	//router.Routes(AuthRouter)      // userShouldNotBeAuthorize, AuthRouter
	routers.ProtectedRouter(app) // checkJwt, authorizeUser, ProtectedRouter

	listingAddress := fmt.Sprintf("localhost:%d", config.Config.Port)
	app.Run(listingAddress)
}

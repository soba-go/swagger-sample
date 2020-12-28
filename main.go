package main

import (
	"restapi/controller"

	"restapi/infrastructure"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	infrastructure.SwaggerStartup()
	ginGonicEngine := gin.Default()
	v1Prefix := "api/v1"
	v1 := ginGonicEngine.Group(v1Prefix)
	{
		ping := v1.Group("/ping")
		{
			ping.GET("", controller.PingPong)
		}
	}
	ginGonicEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginGonicEngine.Run()
}

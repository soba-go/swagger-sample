package infrastructure

import (
	"restapi/docs"
)

// SwaggerStartup Swagger infos startup
func SwaggerStartup() {
	docs.SwaggerInfo.Title = "Swagger Example API Teste"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

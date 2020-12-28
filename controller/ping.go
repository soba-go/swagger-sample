package controller

import (
	"net/http"

	"restapi/models"

	"github.com/gin-gonic/gin"
)

// PingPong godoc
// @Summary Pings a Server
// @Description ping a server
// @Produce  json
// @Success 200 {object} controller.PingPongResult
// @Router /ping [get]
func PingPong(ctx *gin.Context) {
	var result = models.NewPingPongResult("Pong")
	ctx.JSON(http.StatusOK, result)
}

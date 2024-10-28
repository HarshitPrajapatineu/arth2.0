package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteManager(server *gin.Engine) {

	server.GET("/api/view/:path", GetView)
	server.GET("/", Test)
}

func Test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Helloooo"})
}

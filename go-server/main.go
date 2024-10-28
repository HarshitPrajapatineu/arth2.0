package main

import (
	"arth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RouteManager(server)
	server.Run(":8080") // localhost:8080

}

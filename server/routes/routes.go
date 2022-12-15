package routes

import (
	"saved-pictures-holder/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter - Setup the router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.HelloWorld)

	return r
}

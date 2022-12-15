package routes

import (
	"saved-pictures-holder/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// SetupRouter - Setup the router
func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.HelloWorld)

	r.POST("v1/api/login", controllers.Login(db))
	// r.GET("v1/api/register", controllers.Auth(db))

	return r
}

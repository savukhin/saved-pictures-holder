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
	r.GET("/v1/api/protectedhealth", controllers.HelloWorld)
	r.GET("/v1/api/protected", controllers.Protected(db))

	r.POST("/v1/api/login", controllers.Login(db))
	r.POST("/v1/api/register/", controllers.Register(db))

	return r
}

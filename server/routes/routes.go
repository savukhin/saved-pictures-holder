package routes

import (
	"saved-pictures-holder/controllers"
	"saved-pictures-holder/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// SetupRouter - Setup the router
func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.HelloWorld)
	r.GET("/v1/api/health", controllers.HelloWorld)
	r.GET("/v1/api/protected", controllers.Protected(db))

	r.POST("/v1/api/auth/login", controllers.Login(db))
	r.POST("/v1/api/auth/register/", controllers.Register(db))

	r.DELETE("/v1/api/auth/delete-myself", controllers.DeleteUser(db))

	r.POST("/v1/api/folders/create", middleware.AuthRequired(db), controllers.CreateFolder(db))

	return r
}

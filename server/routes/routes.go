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
	r.GET("/v1/api/folders/get/all", middleware.AuthRequired(db), controllers.GetFolders(db))
	r.GET("/v1/api/folders/get/:id", middleware.AuthRequired(db), controllers.GetFolderByID(db))
	r.DELETE("/v1/api/folders/delete/:id", middleware.AuthRequired(db), controllers.DeleteFolder(db))

	return r
}

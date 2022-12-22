package routes

import (
	"saved-pictures-holder/controllers"
	"saved-pictures-holder/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SetupRouter - Setup the router
func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/", controllers.HelloWorld)
	r.GET("/v1/api/health", controllers.HelloWorld)
	r.GET("/v1/api/protected", controllers.Protected(db))

	r.POST("/v1/api/auth/login", controllers.Login(db))
	r.POST("/v1/api/auth/register/", controllers.Register(db))

	r.DELETE("/v1/api/auth/delete-myself", controllers.DeleteUser(db))

	r.POST("/v1/api/folders/create", middleware.AuthRequired(db), controllers.CreateFolder(db))
	r.GET("/v1/api/folders/get/all", middleware.AuthRequired(db), controllers.GetFolders(db))
	r.GET("/v1/api/folders/get/:id", middleware.AuthRequired(db), controllers.GetFolderByID(db))
	r.PUT("/v1/api/folders/update/:id", middleware.AuthRequired(db), controllers.UpdateFolder(db))
	r.DELETE("/v1/api/folders/delete/:id", middleware.AuthRequired(db), controllers.DeleteFolder(db))

	r.POST("/v1/api/folders/:id/create-picture/", middleware.AuthRequired(db), controllers.CreatePicture(db))
	r.GET("/v1/api/folders/:id/pictures/", middleware.AuthRequired(db), controllers.GetPictures(db)) // Params: offset, limit

	r.GET("/v1/api/picture/:id", middleware.AuthRequired(db), controllers.GetPictureInfo(db))
	r.POST("/v1/api/picture/:id/update", middleware.AuthRequired(db), controllers.UpdatePicture(db))

	return r
}

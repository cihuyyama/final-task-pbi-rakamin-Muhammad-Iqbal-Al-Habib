package router

import (
	"task-5-pbi-btpns-habib/controllers"
	"task-5-pbi-btpns-habib/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	{
		public.POST("/users/login", controllers.Login)
		public.POST("/users/register", controllers.CreateUser)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.Authenticate())
	{
		protected.GET("/users/:id", controllers.GetUserByID)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)
		protected.GET("/photos", controllers.GetAllPhotos)
		protected.GET("/photos/:id", controllers.GetPhotoByID)
		protected.POST("/photos", controllers.CreatePhoto)
		protected.PUT("/photos/:id", controllers.UpdatePhoto)
		protected.DELETE("/photos/:id", controllers.DeletePhoto)
	}

	return r
}

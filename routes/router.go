package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/laironacosta/ms-gin-go/controllers"
)

func NewRouter() *gin.Engine {
	//create a default router with default middleware
	r := gin.Default()

	basePath := r.Group("/ms-gin-go")

	basePath.GET("/health", controllers.Health)

	users := basePath.Group("/users")
	{
		users.POST("/", controllers.Create)
		users.GET("/:email", controllers.GetByEmail)
		users.PUT("/:email", controllers.UpdateByEmail)
		users.PATCH("/:email", controllers.UpdateByEmail)
		users.DELETE("/:email", controllers.DeleteByEmail)
	}

	return r
}

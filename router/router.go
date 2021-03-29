package router

import (
	"github.com/gin-gonic/gin"
	"github.com/laironacosta/ms-gin-go/controllers"
)

type Router struct {
	server         *gin.Engine
	userController controllers.UserControllerInterface
}

func NewRouter(server *gin.Engine, userController controllers.UserControllerInterface) *Router {
	return &Router{
		server,
		userController,
	}
}

func (r *Router) Init() {
	//create a default router with default middleware
	basePath := r.server.Group("/ms-gin-go")

	basePath.GET("/health", controllers.Health)

	users := basePath.Group("/users")
	{
		users.POST("/", r.userController.Create)
		users.GET("/:email", r.userController.GetByEmail)
		users.PUT("/:email", r.userController.UpdateByEmail)
		users.PATCH("/:email", r.userController.UpdateByEmail)
		users.DELETE("/:email", r.userController.DeleteByEmail)
	}
}

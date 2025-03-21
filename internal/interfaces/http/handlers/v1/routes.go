package v1

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, userHandler *UserHandler) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users", userHandler.GetUsers)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.RemoveUser)
	}
}

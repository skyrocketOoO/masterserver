package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/masterserver/internal/delivery/rest"
)

func Binding(r *gin.Engine, d *rest.RestDelivery) {
	r.GET("/ping", d.Ping)
	r.GET("/healthy", d.Healthy)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userR := r.Group("/users")
	{
		userR.GET("", d.GetUsers)
		userR.GET("/:id", d.GetUser)
		userR.POST("", d.CreateUser)
		userR.PUT("/:id", d.UpdateUser)
		userR.PUT("", d.UpdateUsers)
		userR.DELETE("/:id", d.DeleteUser)
		userR.DELETE("", d.DeleteUsers)
	}
}

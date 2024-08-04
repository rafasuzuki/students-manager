package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rafasuzuki/students-manager/handlers"
	"github.com/rafasuzuki/students-manager/middleware"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/students", handlers.Create())
		v1.POST("/login", handlers.GetLogin())
		v1.GET("/signin", middleware.RequireAuth(), handlers.SignIn())
	}
}

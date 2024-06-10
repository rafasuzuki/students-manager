package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rafasuzuki/students-manager/handlers"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/students", handlers.Create())
	}
}

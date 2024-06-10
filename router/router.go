package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rafasuzuki/students-manager/handlers"
)

func Initialize() {
	r := gin.Default()

	r.POST("/students", handlers.Create())

	log.Fatal(r.Run(":8080"))
}

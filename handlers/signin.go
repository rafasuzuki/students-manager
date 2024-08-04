package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"message": "Aula agendada com sucesso"})
	}
}

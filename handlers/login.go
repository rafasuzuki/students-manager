package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafasuzuki/students-manager/models"
)

func GetLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login models.Login
		err := c.ShouldBindJSON(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		login_db, err := models.GetLogin(login.Email)
		fmt.Println(login_db)
		// fazer magica de gerar token e salvar no redis. 
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Login realizado com sucesso"})
	}
}

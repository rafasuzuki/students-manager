package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafasuzuki/students-manager/models"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context){
		var student models.Student
		err := c.ShouldBindJSON(&student)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := models.Insert(student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		   }
		c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Aluno inserido com sucesso! Id: %d", id)})
	}
}

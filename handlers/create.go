package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafasuzuki/students-manager/models"
)

// func Create(w http.ResponseWriter, r *http.Request) {
// 	var student models.Student

// 	err := json.NewDecoder(r.Body).Decode(&student)
// 	if err != nil {
// 		log.Printf("Erro ao fazer decode do Json: %v", err)
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	id, err := models.Insert(student)

// 	var resp map[string]any

// 	if err != nil {
// 		resp = map[string]any{
// 			"status": 400,
// 			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
// 		}
// 	} else {
// 		resp = map[string]any{
// 			"status": 200,
// 			"Message": fmt.Sprintf("Aluno inserido com sucesso! Id: %d", id),
// 		}
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(resp)

// }

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

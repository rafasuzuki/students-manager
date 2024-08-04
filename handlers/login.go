package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rafasuzuki/students-manager/models"
)

func GetLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.Login
		//pegando dados da requisicao
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//procurando pelo request do usuario
		db_data, err := models.GetLogin(body.Email)
		fmt.Println(body)
		fmt.Println(db_data)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//comparar password enviado com o do db
		//TODO: Melhorar comparacao de senhas
		if body.Senha != db_data.Senha {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": "Email ou Senha invalidos",
			})
			return
		}

		// fazer magica do jwt
		token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
			"user": db_data.ID,
			"exp":  time.Now().Add(time.Hour).Unix(),
		})
		//gerando o token encriptado usando o secret
		//TODO: verificar criacao do token com string secret
		tokenString, err := token.SigningString()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": "Nao foi possivel gerar o token",
			})
			return
		}
		//mandar o Token de volta
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "Login realizado com sucesso"})
	}
}

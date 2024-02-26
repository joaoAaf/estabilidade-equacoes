package controler

import (
	"estabilidade-equacoes/input"
	"estabilidade-equacoes/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InputCalc(c *gin.Context) {
	var input input.Input
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}
	result := service.ValidateEquation(input)
	c.IndentedJSON(http.StatusCreated, result)
}

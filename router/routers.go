package router

import (
	"estabilidade-equacoes/controler"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.POST("/calc", controler.InputCalc)
}

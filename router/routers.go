package router

import (
	"estabilidade-equacoes/controler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
	}))
	router.POST("/calc", controler.InputCalc)
	router.Run(":8080")
}

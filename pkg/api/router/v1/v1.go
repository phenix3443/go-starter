package router

import (
	"github.com/gin-gonic/gin"

	"github.com/phenix3443/go-starter/pkg/api/handler"
)

func registerV1Routes(rg *gin.RouterGroup) {
	rg.GET("/person/:name", handler.GetPerson)
	rg.POST("/person/:name", handler.SavePerson)
}

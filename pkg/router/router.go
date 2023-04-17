package router

import (
	"homework/homework_2/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.StaticFile("/", "web/index.html")
	r.GET("/images/:id", api.ReadImage)
	r.POST("/images", api.CreateImage)
	r.PUT("/images/:id", api.UpdateImage)
	r.DELETE("/images/:id", api.DeleteImage)

	return r
}
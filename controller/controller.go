package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

func Run() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/rss2json", rss2json)
		v1.GET("/ip", ip)
		v1.GET("/time", getTime)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(ctx *gin.Context) {
		ctx.Header("Location", "/swagger/index.html")
		ctx.JSON(http.StatusFound, nil)
		ctx.Abort()
		return
	})
	if err := r.Run(":4000"); err != nil {
		panic(err)
	}
}

package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kordondev/go-gin/controller"
	"github.com/kordondev/go-gin/middlewares"
	"github.com/kordondev/go-gin/service"
	ginDump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), ginDump.Dump(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.GetAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Video Input is Valid!"})
		}
	})

	server.Run(":8888")
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kordondev/go-gin/entity"
	"github.com/kordondev/go-gin/service"
)

type VideoController interface {
	GetAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (controller *videoController) GetAll() []entity.Video {
	return controller.service.GetAll()
}

func (controller *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return controller.service.Save(video)
}

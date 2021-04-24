package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kordondev/go-gin/entity"
	"github.com/kordondev/go-gin/service"
	"github.com/kordondev/go-gin/validators"
)

type VideoController interface {
	GetAll() []entity.Video
	Save(ctx *gin.Context) error
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}

func (controller *videoController) GetAll() []entity.Video {
	return controller.service.GetAll()
}

func (controller *videoController) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}

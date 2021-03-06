package interfaces

import (
	"go-DDD/application"
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
	"go-DDD/dto/restdto"

	"github.com/gin-gonic/gin"
)

type LastestNews struct {
	aln application.LastestNewApp
}

func NewLastestNewsController(ln repository.AnnouncementRepo) *LastestNews {
	return &LastestNews{
		aln: *application.CreateNewLastestNewApp(ln),
	}
}

func (ln *LastestNews) GetNewDetail(ctx *gin.Context) {
	result := restdto.Success()
	id := ctx.Query("Id")
	result.Data = ln.aln.GetNewDetail(id)

	ctx.JSON(200, &result)
}

func (ln *LastestNews) CreateNew(ctx *gin.Context) {
	var data entity.Announcement
	ctx.ShouldBindJSON(&data)
	err := ln.aln.CreateNew(data)
	if err != nil {
		ctx.JSON(500, restdto.Fail500())
	}

	ctx.JSON(200, restdto.Success())
}

func (ln *LastestNews) DeleteNew(ctx *gin.Context) {
	var data map[string]string
	ctx.ShouldBindJSON(&data)
	if data["Id"] == "" {
		ctx.JSON(403, restdto.Fail403())
	} else {
		ln.aln.DeleteNew(data["Id"])
		ctx.JSON(200, restdto.Success())
	}
}

func (ln *LastestNews) GetNewList(ctx *gin.Context) {
	result := restdto.Success()
	result.Data = ln.aln.GetListNew()
	ctx.JSON(200, &result)
}

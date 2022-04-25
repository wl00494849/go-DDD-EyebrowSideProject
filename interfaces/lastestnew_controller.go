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

func NewLastestNewsController(ln repository.LatestNewsRepo) *LastestNews {
	return &LastestNews{
		aln: *application.CreateNewLastestNewApp(ln),
	}
}

func (ln *LastestNews) GetNewDetail(ctx *gin.Context) {
	id := ctx.Query("NewID")
	data := ln.aln.GetNewDetail(id)
	ctx.JSON(200, data)
}

func (ln *LastestNews) CreateNew(ctx *gin.Context) {
	var data entity.LatestNews
	ctx.ShouldBindJSON(&data)
	file, _ := ctx.FormFile("Image")
	err := ln.aln.CreateNew(data)
	if err != nil {
		ctx.JSON(500, &restdto.Result{
			IsSuccess: false,
			Msg:       "CreateNew Fail",
		})
	}
	ctx.SaveUploadedFile(file, "/asset/New/"+file.Filename)

	ctx.JSON(200, restdto.Success())
}

func (ln *LastestNews) DeleteNew(ctx *gin.Context) {
	var data map[string]string
	ctx.ShouldBindJSON(&data)
	if data["NewID"] == "" {
		ctx.JSON(200, &restdto.Result{
			IsSuccess: false,
			Msg:       "Not have ID",
		})
	} else {
		ln.aln.DeleteNew(data["NewID"])
		ctx.JSON(200, restdto.Success())
	}
}

func (ln *LastestNews) GetNewList(ctx *gin.Context) {
	result := ln.aln.GetListNew()
	ctx.JSON(200, result)
}

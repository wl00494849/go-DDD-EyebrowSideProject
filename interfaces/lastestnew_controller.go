package interfaces

import (
	"go-DDD/application"
	"go-DDD/domain/entity"

	"github.com/gin-gonic/gin"
)

type LastestNews struct {
	aln application.LastestNewAppInterFace
}

func NewLastestNewsController(ln application.LastestNewAppInterFace) *LastestNews {
	return &LastestNews{
		aln: ln,
	}
}

func (ln *LastestNews) GetNewDetail(ctx *gin.Context) {
	id := ctx.Query("id")
	data := ln.aln.GetNewDetail(id)
	ctx.JSON(200, data)
}

func (ln *LastestNews) SetNew(ctx *gin.Context) {
	var data entity.LatestNews
	ctx.ShouldBindJSON(&data)
	err := ln.aln.SetNew(&data)
	if err != nil {
		ctx.JSON(500, "Set fail")
	}
	ctx.JSON(200, "sucess")
}

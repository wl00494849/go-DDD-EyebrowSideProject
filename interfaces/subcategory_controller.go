package interfaces

import (
	"go-DDD/application"
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
	"go-DDD/dto/restdto"

	"github.com/gin-gonic/gin"
)

type SubCategory struct {
	as application.SubcategoryApp
}

func NewSubcategoryController(rs repository.SubcategoryRepo) *SubCategory {
	return &SubCategory{as: *application.CreateSysCodeApp(rs)}
}

func (s *SubCategory) CreateSubCategory(ctx *gin.Context) {
	var data entity.Subcategory
	ctx.BindJSON(&data)
	err := s.as.CreateSubCategory(&data)
	if err != nil {
		ctx.JSON(500, restdto.Fail500())
		return
	}
	ctx.JSON(200, restdto.Success())
}

func (s *SubCategory) GetSubCategorys(ctx *gin.Context) {
	result := restdto.Success()
	code := ctx.Query("Id")
	result.Data = s.as.GetSubCategorys(code)
	ctx.JSON(200, result)
}

func (s *SubCategory) DeleteSubCategory(ctx *gin.Context) {
	var data map[string]string
	ctx.BindJSON(&data)

	if data["Id"] == "" {
		ctx.JSON(500, restdto.Fail500())
		return
	}

	if err := s.as.DeleteSubCategory(data["Id"]); err != nil {
		ctx.JSON(500, restdto.Fail500())
		return
	}

	ctx.JSON(200, restdto.Success())
}

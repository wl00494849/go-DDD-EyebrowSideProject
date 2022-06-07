package main

import (
	"flag"
	"fmt"
	"go-DDD/infrastruture/persistence"
	"go-DDD/interfaces"
	"go-DDD/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env got")
	}
}

func main() {
	var port string

	flag.StringVar(&port, "p", ":8080", "port")
	if p := os.Getenv("PORT"); len(p) != 0 {
		port = ":" + p
	}
	//database
	dbUser := os.Getenv("MYSQL_USER")
	dbPwd := os.Getenv("MYSQL_PASSWORD")
	dbAddr := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("MYSQL_DATABASE")

	servers, err := persistence.NewRepositories(dbUser, dbPwd, dbAddr, dbName)
	if err != nil {
		fmt.Println(err)
	}

	defer servers.Close()
	newController := interfaces.NewLastestNewsController(servers.AnnouncementRepo)
	subcategoryController := interfaces.NewSubcategoryController(servers.SubCategoryRepo)
	//app
	app := gin.Default()
	//middle
	app.Use(middleware.CORSMiddleware())
	//route
	news := app.Group("/News")
	{
		news.GET("/GetDetail", newController.GetNewDetail)
		news.POST("/Create", newController.CreateNew)
		news.POST("/Delete", newController.DeleteNew)
		news.GET("/GetList", newController.GetNewList)
	}
	subcategory := app.Group("/Subcategory")
	{
		subcategory.POST("/Create", subcategoryController.CreateSubCategory)
		subcategory.POST("/Delete", subcategoryController.DeleteSubCategory)
		subcategory.GET("/GetList", subcategoryController.GetSubCategorys)
	}

	app.Run(port)
}

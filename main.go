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
	//app
	app := gin.Default()
	//middle
	app.Use(middleware.CORSMiddleware())
	//route
	news := app.Group("/News")
	{
		news.GET("/GetNewDetail", newController.GetNewDetail)
		news.POST("/CreateNew", newController.CreateNew)
		news.POST("/DeleteNew", newController.DeleteNew)
		news.GET("/GetNewList", newController.GetNewList)
	}

	app.Run(port)
}

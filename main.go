package main

import (
	"flag"
	"fmt"
	"go-DDD/controller"
	"go-DDD/infrastruture/persistence"
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

	newController := controller.NewLastestNewsController(servers.LatestNewsRepo)

	//app
	app := gin.Default()
	//middle
	app.Use(middleware.CORSMiddleware())
	//route
	app.GET("/GetNewDetail", newController.GetNewDetail)
	app.POST("/SetNew", newController.SetNew)
	app.Run(port)
}

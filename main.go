package main

import (
	"flag"
	"fmt"
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

	app := gin.Default()

	app.Run(port)
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/otaviobernardes/api-go-gin/src/external/database"
	"github.com/otaviobernardes/api-go-gin/src/routes"
)

func main() {
	r := gin.Default()
	routes.HandleRequest(r)
	database.New()
	r.Run(":3000")
}
